// +build wireinject

package main

import (
	"github.com/blackhorseya/ip-rate-limit/internal/app"
	"github.com/blackhorseya/ip-rate-limit/internal/app/apis"
	"github.com/blackhorseya/ip-rate-limit/internal/app/biz"
	"github.com/blackhorseya/ip-rate-limit/pkg/config"
	"github.com/blackhorseya/ip-rate-limit/pkg/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	app.ProviderSet,
	config.ProviderSet,
	http.ProviderSet,
	apis.ProviderSet,
	biz.ProviderSet,
)

// CreateInjector serve caller to create an injector
func CreateInjector(path string) (*app.Injector, error) {
	panic(wire.Build(providerSet))
}
