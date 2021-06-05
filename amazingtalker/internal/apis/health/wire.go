// +build wireinject

package health

import (
	"github.com/blackhorseya/amazingtalker/internal/biz/health"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(biz health.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
