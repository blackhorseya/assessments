package health

import (
	"github.com/blackhorseya/amazingtalker/internal/biz/health/repo"
	"github.com/blackhorseya/amazingtalker/internal/pkg/contextx"
)

type impl struct {
	repo repo.IRepo
}

// NewImpl serve caller to create an IBiz
func NewImpl(repo repo.IRepo) IBiz {
	return &impl{repo: repo}
}

func (i *impl) Readiness(ctx contextx.Contextx) (bool, error) {
	ok, err := i.repo.Ping(ctx)
	if err != nil {
		ctx.WithError(err).Error("ping db is failure")
		return false, err
	}

	return ok, nil
}

func (i *impl) Liveness(ctx contextx.Contextx) (bool, error) {
	ok, err := i.repo.Ping(ctx)
	if err != nil {
		ctx.WithError(err).Error("ping db is failure")
		return false, err
	}

	return ok, nil
}
