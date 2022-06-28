package service

import (
	"github.com/hunkevych-philip/mono-app/pkg/service/excel"
	"github.com/hunkevych-philip/mono-app/pkg/service/mono"
	"github.com/hunkevych-philip/mono-app/pkg/types"
	"time"
)

//Mono is a custom service for retrieving data from the monobank using an API client
type Mono interface {
	ProcessStatement(token, account string, startDate time.Time) (*types.Statement, error)
}

type Excel interface {
	GenerateSheetForStatement(statement *types.Statement) error
}

type ServicesImpl struct {
	Mono  Mono
	Excel Excel
}

func NewService() (*ServicesImpl, error) {
	excelService := excel.NewExcelService()
	monoService, err := mono.NewMonoService()
	if err != nil {
		return nil, err
	}

	return &ServicesImpl{
		Mono:  monoService,
		Excel: excelService,
	}, nil
}
