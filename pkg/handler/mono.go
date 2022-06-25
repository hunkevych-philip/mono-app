package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

const (
	HeaderKeyXToken string = "X-Token"
	ParamKeyAccount string = "account"
	ParamKeyFrom    string = "from"
	ParamKeyTo      string = "to"
)

func (h *Handler) getClientStatement(c *gin.Context) {
	var (
		token     = c.GetHeader(HeaderKeyXToken)
		account   = c.Param(ParamKeyAccount)
		startDate = c.Param(ParamKeyFrom) // we expect it to be a full date
		_         = c.Param(ParamKeyTo)   // empty is ok, a format the same as for startDate

		RFC3339noTimeLayout = time.RFC3339[:strings.Index(time.RFC3339, "T")]
	)

	timeT, err := time.Parse(RFC3339noTimeLayout, "2020-04-14")
	if err != nil {
		// TODO: Handler error
		logrus.Errorf("Failed to parse date: %s", err.Error())

		return
	}

	if len(token) == 0 {
		// TODO: Handle error
		logrus.Error("X-Token header is missing")

		return
	}
	if len(account) == 0 {
		account = "0" // default
	}
	if len(startDate) == 0 {
		// TODO: Handle error
		logrus.Errorf("%q parameter cannot be empty", ParamKeyFrom)

		return
	}

	// TODO: Make an API call to a monobank here
	fmt.Println(timeT)
}

func (h *Handler) getClientInfo(c *gin.Context) {
}
