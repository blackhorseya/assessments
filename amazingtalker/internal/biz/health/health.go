package health

import (
	"github.com/blackhorseya/amazingtalker/internal/biz/health/repo"
	"github.com/blackhorseya/amazingtalker/internal/pkg/contextx"
	"github.com/google/wire"
)

// IBiz declare business service function
type IBiz interface {
	// Readiness to handle application has been ready
	Readiness(ctx contextx.Contextx) (bool, error)

	// Liveness to handle application was alive
	Liveness(ctx contextx.Contextx) (bool, error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
