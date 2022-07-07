package mono_client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hunkevych-philip/mono-app/pkg/types"
	"io/ioutil"
	"net/http"
	"time"
)

type MonoClient struct {
	HttpClient *http.Client
	BaseUrl    string
}

func NewMonoClient() *MonoClient {
	return &MonoClient{
		HttpClient: new(http.Client),
		// TODO: We can pass BaseUrl as an input
		BaseUrl: "https://api.monobank.ua",
	}
}

func (c *MonoClient) GetStatement(token, account string, startDate time.Time) (*types.Statement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// https://api.monobank.ua/personal/statement/{account}/{from}/{to}
	url := c.BaseUrl + fmt.Sprintf("/personal/statement/%s/%d}", account, startDate.Unix())
	fmt.Println(url)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add(types.HeaderKeyXToken, token)

	do, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if do.StatusCode != http.StatusOK {
		// TODO: Handle response error
		return nil, fmt.Errorf("mono API returned an unexpected status code: %d", do.StatusCode)
	}

	all, err := ioutil.ReadAll(do.Body)
	if err != nil {
		return nil, err
	}

	res := &types.Statement{
		StatementRecords: make([]*types.StatementRecord, 0),
	}
	if err := json.Unmarshal(all, &res.StatementRecords); err != nil {
		return nil, err
	}

	return res, nil
}
