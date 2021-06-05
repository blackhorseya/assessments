// +build wireinject

package repo

import (
	"github.com/blackhorseya/amazingtalker/internal/pkg/config"
	"github.com/blackhorseya/amazingtalker/internal/pkg/databases"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	config.ProviderSet,
	databases.ProviderSet,
	NewImpl,
)

// CreateRepo serve caller to create an IRepo
func CreateRepo(path string) (IRepo, error) {
	panic(wire.Build(testProviderSet))
}