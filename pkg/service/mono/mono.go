package mono

import (
	"fmt"
	"github.com/hunkevych-philip/mono-app/pkg/service/mono/mono_client"
	"github.com/hunkevych-philip/mono-app/pkg/types"
	"github.com/sirupsen/logrus"
	"time"
)

type MonoClient interface {
	GetStatement(token, account string, startDate time.Time) (*types.Statement, error)
}

type MonoService struct {
	MonoClient MonoClient
}

func NewMonoService() (*MonoService, error) {
	monoClient, err := mono_client.NewMonoClient()
	if err != nil {
		return nil, err
	}

	return &MonoService{
		MonoClient: monoClient,
	}, nil
}

func (s *MonoService) ProcessStatement(token, account string, startDate time.Time) (*types.Statement, error) {
	// max possible value is 31 days + 1 hour
	if time.Now().UnixNano()-startDate.UnixNano() > int64(time.Hour*24*31+time.Hour) {
		err := fmt.Errorf("a start date should be within 31 day range from now. You entered: %s", startDate)
		logrus.Error(err)

		return nil, err
	}

	statement, err := s.MonoClient.GetStatement(token, account, startDate)
	if err != nil {
		// TODO: Implement error handling recursively
		logrus.Error(err)

		return nil, err
	}

	// TODO: Export data to some file

	return statement, nil
}
