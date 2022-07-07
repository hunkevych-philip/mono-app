package service

import (
	"github.com/hunkevych-philip/mono-app/pkg/types"
)

//Mono is a custom service for retrieving data from the monobank using an API client
type Mono interface {
	GetStatement(token, account string, from string) (*types.Statement, error)
}

type Excel interface {
	GenerateSheetForStatement(statement *types.Statement) error
}

type ServicesImpl struct {
	Mono  Mono
	Excel Excel
}

func NewService(mono Mono, excel Excel) *ServicesImpl {
	return &ServicesImpl{
		Mono:  mono,
		Excel: excel,
	}
}
