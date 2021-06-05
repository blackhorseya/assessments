package apis

import (
	"time"

	// import swagger docs
	_ "github.com/blackhorseya/amazingtalker/internal/apis/docs"
	"github.com/blackhorseya/amazingtalker/internal/apis/health"
	"github.com/blackhorseya/amazingtalker/internal/apis/tutors"
	"github.com/blackhorseya/amazingtalker/internal/pkg/config"
	"github.com/blackhorseya/amazingtalker/internal/pkg/transports/http"
	"github.com/blackhorseya/amazingtalker/internal/pkg/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/patrickmn/go-cache"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn(cfg *config.Config, health health.IHandler, tutorsHandler tutors.IHandler) http.InitHandlers {
	store := cache.New(time.Duration(cfg.HTTP.CacheExpireMin)*time.Minute, time.Duration(cfg.HTTP.CacheExpireMin)*time.Minute)

	return func(r *gin.Engine) {

		api := r.Group("api")
		{
			api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

			api.GET("readiness", health.Readiness)
			api.GET("liveness", health.Liveness)

			api.GET("tutors/:slug", middlewares.CacheMiddleware(store), tutorsHandler.ListByLangSlug)
			api.GET("tutor/:slug", middlewares.CacheMiddleware(store), tutorsHandler.GetInfoBySlug)
		}
	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	tutors.ProviderSet,
	CreateInitHandlerFn,
)
