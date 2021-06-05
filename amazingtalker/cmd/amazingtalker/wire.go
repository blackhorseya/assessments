// +build wireinject

package main

import (
	"github.com/blackhorseya/amazingtalker/internal/apis"
	"github.com/blackhorseya/amazingtalker/internal/biz"
	"github.com/blackhorseya/amazingtalker/internal/pkg/app"
	"github.com/blackhorseya/amazingtalker/internal/pkg/config"
	"github.com/blackhorseya/amazingtalker/internal/pkg/databases"
	"github.com/blackhorseya/amazingtalker/internal/pkg/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	app.ProviderSet,
	config.ProviderSet,
	http.ProviderSet,
	databases.ProviderSet,
	apis.ProviderSet,
	biz.ProviderSet,
)

// CreateInjector serve caller to create an app
func CreateInjector(path string) (*app.Injector, error) {
	panic(wire.Build(providerSet))
}
