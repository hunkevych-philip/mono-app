package service

import (
	"github.com/hunkevych-philip/mono-app/pkg/service/mono"
	"github.com/hunkevych-philip/mono-app/pkg/types"
	"time"
)

//Mono is a custom service for retrieving data from the monobank using an API client
type Mono interface {
	ProcessStatement(token, account string, startDate time.Time) ([]*types.Statement, error)
}

type ServicesImpl struct {
	Mono Mono
}

func NewService() (*ServicesImpl, error) {
	monoService, err := mono.NewMonoService()
	if err != nil {
		return nil, err
	}

	return &ServicesImpl{
		Mono: monoService,
	}, nil
}
