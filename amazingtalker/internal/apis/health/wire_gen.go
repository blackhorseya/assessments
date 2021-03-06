// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package health

import (
	"github.com/blackhorseya/amazingtalker/internal/biz/health"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateIHandler(biz health.IBiz) (IHandler, error) {
	iHandler := NewImpl(biz)
	return iHandler, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)
