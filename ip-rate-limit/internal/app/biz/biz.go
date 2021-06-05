package biz

import (
	"github.com/blackhorseya/ip-rate-limit/internal/app/biz/counter"
	"github.com/google/wire"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	counter.ProviderSet,
)
