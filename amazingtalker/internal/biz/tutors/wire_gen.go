// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package tutors

import (
	"github.com/blackhorseya/amazingtalker/internal/biz/tutors/repo"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateIBiz(repo2 repo.IRepo) (IBiz, error) {
	iBiz := NewImpl(repo2)
	return iBiz, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)