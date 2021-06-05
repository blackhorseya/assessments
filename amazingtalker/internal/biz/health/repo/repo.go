package repo

import (
	"github.com/blackhorseya/amazingtalker/internal/pkg/contextx"
	"github.com/google/wire"
)

// IRepo declare repository function
type IRepo interface {
	Ping(ctx contextx.Contextx) (bool, error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
