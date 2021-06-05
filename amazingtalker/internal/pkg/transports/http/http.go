package http

import (
	"github.com/blackhorseya/amazingtalker/internal/pkg/config"
	"github.com/blackhorseya/amazingtalker/internal/pkg/transports/http/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// InitHandlers define register handler
type InitHandlers func(r *gin.Engine)

// NewGinEngine serve caller to create a gin.Engine
func NewGinEngine(cfg *config.Config, init InitHandlers) *gin.Engine {
	gin.SetMode(cfg.HTTP.Mode)

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(middlewares.ContextMiddleware())
	r.Use(middlewares.LoggerMiddleware())

	init(r)

	return r
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewGinEngine)
