package service

import "github.com/HunkevychPhilip/todo/pkg/service/mono"

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
