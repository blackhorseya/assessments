package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type impl struct {
}

// NewImpl serve caller to create an IHandler
func NewImpl() IHandler {
	return &impl{}
}

// @Summary Readiness
// @Description Show application was ready to start accepting traffic
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /api/readiness [get]
func (i *impl) Readiness(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// @Summary Liveness
// @Description to know when to restart an application
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /api/liveness [get]
func (i *impl) Liveness(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
