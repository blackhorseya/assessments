// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/blackhorseya/amazingtalker/internal/apis"
	health2 "github.com/blackhorseya/amazingtalker/internal/apis/health"
	tutors2 "github.com/blackhorseya/amazingtalker/internal/apis/tutors"
	"github.com/blackhorseya/amazingtalker/internal/biz"
	"github.com/blackhorseya/amazingtalker/internal/biz/health"
	"github.com/blackhorseya/amazingtalker/internal/biz/health/repo"
	"github.com/blackhorseya/amazingtalker/internal/biz/tutors"
	repo2 "github.com/blackhorseya/amazingtalker/internal/biz/tutors/repo"
	"github.com/blackhorseya/amazingtalker/internal/pkg/app"
	"github.com/blackhorseya/amazingtalker/internal/pkg/config"
	"github.com/blackhorseya/amazingtalker/internal/pkg/databases"
	"github.com/blackhorseya/amazingtalker/internal/pkg/transports/http"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateInjector(path2 string) (*app.Injector, error) {
	configConfig, err := config.NewConfig(path2)
	if err != nil {
		return nil, err
	}
	db, err := databases.NewMariaDB(configConfig)
	if err != nil {
		return nil, err
	}
	iRepo := repo.NewImpl(db)
	iBiz := health.NewImpl(iRepo)
	iHandler := health2.NewImpl(iBiz)
	repoIRepo := repo2.NewImpl(db)
	tutorsIBiz := tutors.NewImpl(repoIRepo)
	tutorsIHandler := tutors2.NewImpl(tutorsIBiz)
	initHandlers := apis.CreateInitHandlerFn(configConfig, iHandler, tutorsIHandler)
	engine := http.NewGinEngine(configConfig, initHandlers)
	injector := app.NewInjector(engine, configConfig)
	return injector, nil
}

// wire.go:

var providerSet = wire.NewSet(app.ProviderSet, config.ProviderSet, http.ProviderSet, databases.ProviderSet, apis.ProviderSet, biz.ProviderSet)
