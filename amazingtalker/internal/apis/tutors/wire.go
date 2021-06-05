// +build wireinject

package tutors

import (
	"github.com/blackhorseya/amazingtalker/internal/biz/tutors"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(biz tutors.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
