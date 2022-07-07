package mono

import (
	"github.com/hunkevych-philip/mono-app/pkg/service/mono/mono_client"
	"github.com/hunkevych-philip/mono-app/pkg/types"
	"strings"
	"time"
)

type MonoClient interface {
	GetStatement(token, account string, startDate time.Time) (*types.Statement, error)
}

type MonoService struct {
	MonoClient MonoClient
}

func NewMonoService() *MonoService {
	return &MonoService{
		MonoClient: mono_client.NewMonoClient(),
	}
}

func (s *MonoService) GetStatement(token, account string, fromStr string) (*types.Statement, error) {
	var (
		err error

		fromTime       = time.Time{}
		RFC3339trimmed = time.RFC3339[:strings.Index(time.RFC3339, "T")]
	)

	if fromStr == "" {
		// Put max allowed value (31 days + 1 hour)
		fromTime = time.Now().Add(-time.Hour * 24 * 31)
	} else {
		fromTime, err = time.Parse(RFC3339trimmed, fromStr)
		if err != nil {
			return nil, err
		}
	}

	if account == "" {
		account = "0" // default account
	}

	statement, err := s.MonoClient.GetStatement(token, account, fromTime)
	if err != nil {
		return nil, err
	}

	return statement, nil
}
