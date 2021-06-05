package health

import (
	"fmt"
	"net/http"

	"github.com/blackhorseya/amazingtalker/internal/biz/health"
	"github.com/blackhorseya/amazingtalker/internal/pkg/contextx"
	"github.com/gin-gonic/gin"
)

var (
	errReadiness = fmt.Errorf("readiness is failure")
	errLiveness  = fmt.Errorf("liveness is failure")
)

type impl struct {
	biz health.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(biz health.IBiz) IHandler {
	return &impl{biz: biz}
}

// @Summary Readiness
// @Description Show application was ready to start accepting traffic
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /readiness [get]
func (i *impl) Readiness(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)
	logger := ctx.WithField("func", "readiness")

	ok, err := i.biz.Readiness(ctx)
	if err != nil {
		logger.WithError(err).Error(errReadiness)
		c.String(http.StatusInternalServerError, errReadiness.Error())
		return
	}
	if !ok {
		logger.Error(errReadiness)
		c.String(http.StatusServiceUnavailable, errReadiness.Error())
		return
	}

	c.String(http.StatusOK, "ok")
}

// @Summary Liveness
// @Description to know when to restart an application
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /liveness [get]
func (i *impl) Liveness(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)
	logger := ctx.WithField("func", "liveness")

	ok, err := i.biz.Liveness(ctx)
	if err != nil {
		logger.WithError(err).Error(errLiveness)
		c.String(http.StatusInternalServerError, errLiveness.Error())
		return
	}
	if !ok {
		logger.Error(errLiveness)
		c.String(http.StatusServiceUnavailable, errLiveness.Error())
		return
	}

	c.String(http.StatusOK, "ok")
}
