package service

import "github.com/hunkevych-philip/mono-app/pkg/service/mono"

type Mono interface {
	Create() (int, error)
}

type ServicesImpl struct {
	Mono Mono
}

func NewService() *ServicesImpl {
	return &ServicesImpl{
		Mono: mono.NewMonoService(),
	}
}
