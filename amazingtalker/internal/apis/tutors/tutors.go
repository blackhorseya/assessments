package tutors

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare health api handler
type IHandler interface {
	ListByLangSlug(c *gin.Context)

	GetInfoBySlug(c *gin.Context)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
