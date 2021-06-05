// +build wireinject

package health

import "github.com/google/wire"

var testProviderSet = wire.NewSet(NewImpl)

// CreateHandler serve caller to new an IHandler
func CreateHandler() (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
