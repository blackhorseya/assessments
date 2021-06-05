package http

import (
	"github.com/blackhorseya/ip-rate-limit/pkg/config"
	"github.com/blackhorseya/ip-rate-limit/pkg/transports/http/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewGinEngine serve caller to create a gin.Engine
func NewGinEngine(cfg *config.Config, init InitHandlers) *gin.Engine {
	gin.SetMode(cfg.HTTP.Mode)

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(middlewares.ContextMiddleware())
	r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.IPRateLimitMiddleware(cfg.HTTP.Limit, cfg.HTTP.Burst))

	init(r)

	return r
}
