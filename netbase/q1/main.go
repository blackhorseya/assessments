package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()

	engine.GET("/api/v1/endpoints", func(c *gin.Context) {
		url := c.Query("url")
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		c.String(http.StatusOK, fmt.Sprintf("%x", md5.Sum(bytes)))
	})

	engine.Run()
}
