package tutors

import (
	"net/http"

	"github.com/blackhorseya/amazingtalker/internal/biz/tutors"
	"github.com/blackhorseya/amazingtalker/internal/pkg/contextx"
	"github.com/blackhorseya/amazingtalker/internal/pkg/responses"
	"github.com/gin-gonic/gin"
)

type impl struct {
	biz tutors.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(biz tutors.IBiz) IHandler {
	return &impl{biz: biz}
}

type request struct {
	Slug string `uri:"slug" binding:"required"`
}

// @Summary GetInfoBySlug
// @Description get tutor by slug
// @Tags Tutors
// @Accept application/json
// @Produce application/json
// @Param slug path string true "Tutor slug"
// @Success 200 {object} pb.Tutor
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /tutor/{slug} [get]
func (i *impl) GetInfoBySlug(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)
	logger := ctx.WithField("func", "GetInfoBySlug")

	var req request
	err := c.ShouldBindUri(&req)
	if err != nil {
		logger.WithError(err).Error("bind uri slug is failure")
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	data, err := i.biz.GetInfoBySlug(ctx, req.Slug)
	if err != nil {
		logger.WithError(err).Error("get info by slug is failure")
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	if data == nil {
		logger.Warn("tutor slug is not found")
		c.JSON(http.StatusNotFound, gin.H{"msg": "tutor slug is not found"})
		return
	}

	c.JSON(http.StatusOK, &responses.DataResp{Data: data})
}

// @Summary ListByLangSlug
// @Description get tutor list by language slug
// @Tags Tutors
// @Accept application/json
// @Produce application/json
// @Param slug path string true "Language slug"
// @Success 200 {array} pb.Tutor
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /tutors/{slug} [get]
func (i *impl) ListByLangSlug(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)
	logger := ctx.WithField("func", "ListByLangSlug")

	var req request
	err := c.ShouldBindUri(&req)
	if err != nil {
		logger.WithError(err).Error("bind uri slug is failure")
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	data, err := i.biz.ListByLangSlug(ctx, req.Slug)
	if err != nil {
		logger.WithError(err).Error("list tutors by language's slug is failure")
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	if len(data) == 0 {
		logger.Warn("list tutors by language's slug is not found")
		c.JSON(http.StatusNotFound, gin.H{"msg": "list tutors by language's slug is not found"})
		return
	}

	c.JSON(http.StatusOK, &responses.DataResp{Data: data})
}
