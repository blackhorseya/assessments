package middlewares

import (
	ratex "github.com/blackhorseya/ip-rate-limit/pkg/rate"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	ip *ratex.IP
)

// IPRateLimitMiddleware serve caller to added ip rate limit into gin
func IPRateLimitMiddleware(limit int, burst int) gin.HandlerFunc {
	if ip == nil {
		ip = ratex.NewIP(rate.Limit(limit), burst)
	}

	return func(c *gin.Context) {
		limiter := ip.Limiter(c.ClientIP())
		c.Set("limiter", limiter)

		c.Next()
	}
}
