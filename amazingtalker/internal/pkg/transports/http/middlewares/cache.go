package middlewares

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

type respBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r respBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r respBodyWriter) WriteString(s string) (int, error) {
	r.body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}

func CacheMiddleware(store *cache.Cache) gin.HandlerFunc {
	queue := make(chan bool, 1)

	return func(c *gin.Context) {
		w := &respBodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = w

		path := c.Request.URL.Path
		result, ok := store.Get(path)
		if ok {
			c.Writer.Write(result.([]byte))
			c.AbortWithStatus(http.StatusOK)
			return
		}

		select {
		case queue <- true:
			result, ok := store.Get(path)
			if ok {
				c.Writer.Write(result.([]byte))
				c.AbortWithStatus(http.StatusOK)
				return
			}

			c.Next()

			if !ok && http.StatusOK <= c.Writer.Status() && c.Writer.Status() < http.StatusMultipleChoices {
				store.Set(path, w.body.Bytes(), cache.DefaultExpiration)
			}

			<-queue
		}
	}
}
