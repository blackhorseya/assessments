package tutors

import (
	"fmt"

	"github.com/blackhorseya/amazingtalker/internal/biz/tutors/repo"
	"github.com/blackhorseya/amazingtalker/internal/pkg/contextx"
	"github.com/blackhorseya/amazingtalker/pb"
)

var (
	errGetInfo = fmt.Errorf("get tutor info by tutor's slug is failure")
	errList    = fmt.Errorf("get tutor info by language's slug is failure")
)

type impl struct {
	repo repo.IRepo
}

func NewImpl(repo repo.IRepo) IBiz {
	return &impl{repo: repo}
}

func (i *impl) GetInfoBySlug(ctx contextx.Contextx, slug string) (*pb.Tutor, error) {
	ret, err := i.repo.GetInfoBySlug(ctx, slug)
	if err != nil {
		ctx.WithError(err).Error(errGetInfo)
		return nil, err
	}

	return ret, nil
}

func (i *impl) ListByLangSlug(ctx contextx.Contextx, slug string) ([]*pb.Tutor, error) {
	ret, err := i.repo.ListByLangSlug(ctx, slug)
	if err != nil {
		ctx.WithError(err).Error(errList)
		return nil, err
	}

	return ret, nil
}
