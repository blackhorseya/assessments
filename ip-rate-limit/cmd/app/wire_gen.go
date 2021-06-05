// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/blackhorseya/ip-rate-limit/internal/app"
	"github.com/blackhorseya/ip-rate-limit/internal/app/apis"
	counter2 "github.com/blackhorseya/ip-rate-limit/internal/app/apis/counter"
	"github.com/blackhorseya/ip-rate-limit/internal/app/apis/health"
	"github.com/blackhorseya/ip-rate-limit/internal/app/biz"
	"github.com/blackhorseya/ip-rate-limit/internal/app/biz/counter"
	"github.com/blackhorseya/ip-rate-limit/pkg/config"
	"github.com/blackhorseya/ip-rate-limit/pkg/transports/http"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateInjector(path2 string) (*app.Injector, error) {
	configConfig, err := config.NewConfig(path2)
	if err != nil {
		return nil, err
	}
	iHandler := health.NewImpl()
	iBiz := counter.NewImpl()
	counterIHandler := counter2.NewImpl(iBiz)
	initHandlers := apis.CreateInitHandlerFn(iHandler, counterIHandler)
	engine := http.NewGinEngine(configConfig, initHandlers)
	injector := app.NewInjector(engine, configConfig)
	return injector, nil
}

// wire.go:

var providerSet = wire.NewSet(app.ProviderSet, config.ProviderSet, http.ProviderSet, apis.ProviderSet, biz.ProviderSet)