package apis

import (
	"github.com/blackhorseya/ip-rate-limit/internal/app/apis/counter"
	// import swagger docs
	_ "github.com/blackhorseya/ip-rate-limit/internal/app/apis/docs"
	"github.com/blackhorseya/ip-rate-limit/internal/app/apis/health"
	"github.com/blackhorseya/ip-rate-limit/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn(health health.IHandler, counter counter.IHandler) http.InitHandlers {
	return func(r *gin.Engine) {
		r.GET("", counter.Count)

		api := r.Group("/api")
		{
			api.GET("readiness", health.Readiness)
			api.GET("liveness", health.Liveness)

			api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	counter.ProviderSet,
	CreateInitHandlerFn,
)
