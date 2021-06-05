package counter

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/ip-rate-limit/internal/app/biz/counter"
	"github.com/blackhorseya/ip-rate-limit/pkg/contextx"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type impl struct {
	biz counter.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(biz counter.IBiz) IHandler {
	return &impl{biz: biz}
}

// @Summary Count
// @Description Count request times
// @Tags Counter
// @Accept application/json
// @Accept text/plain
// @Produce application/json
// @Produce text/plain
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router / [get]
func (i *impl) Count(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)
	logger := ctx.WithField("func", "counter count")
	limiter := c.MustGet("limiter").(*rate.Limiter)

	count, err := i.biz.Count(ctx, c.ClientIP(), limiter)
	if err != nil {
		logger.WithField("err", err).Error("Error")
		c.String(http.StatusInternalServerError, "Error")
		return
	}

	c.String(http.StatusOK, strconv.Itoa(count))
}
