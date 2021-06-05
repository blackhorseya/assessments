// +build wireinject

package counter

import (
	"github.com/blackhorseya/ip-rate-limit/internal/app/biz/counter"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateHandler serve caller to new an IHandler
func CreateHandler(biz counter.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
