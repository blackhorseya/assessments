package tutors

import (
	"github.com/blackhorseya/amazingtalker/internal/biz/tutors/repo"
	"github.com/blackhorseya/amazingtalker/internal/pkg/contextx"
	"github.com/blackhorseya/amazingtalker/pb"
	"github.com/google/wire"
)

// IBiz describe business service function
type IBiz interface {
	GetInfoBySlug(ctx contextx.Contextx, slug string) (*pb.Tutor, error)

	ListByLangSlug(ctx contextx.Contextx, slug string) ([]*pb.Tutor, error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
