package app

import (
	"github.com/blackhorseya/amazingtalker/internal/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// Injector declare app something
type Injector struct {
	Engine *gin.Engine
	C      *config.Config
}


// NewInjector serve caller to create an app
func NewInjector(r *gin.Engine, c *config.Config) *Injector {
	return &Injector{Engine: r, C: c}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewInjector)
