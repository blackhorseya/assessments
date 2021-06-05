package counter

import (
	"github.com/blackhorseya/ip-rate-limit/pkg/contextx"
	"github.com/google/wire"
	"golang.org/x/time/rate"
)

type IBiz interface {
	Count(ctx contextx.Contextx, ip string, limiter *rate.Limiter) (int, error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
