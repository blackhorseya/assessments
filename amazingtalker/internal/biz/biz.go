package biz

import (
	"github.com/blackhorseya/amazingtalker/internal/biz/health"
	"github.com/blackhorseya/amazingtalker/internal/biz/tutors"
	"github.com/google/wire"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	tutors.ProviderSet,
)
