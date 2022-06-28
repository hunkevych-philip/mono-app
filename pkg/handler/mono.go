package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hunkevych-philip/mono-app/pkg/types"
	"github.com/hunkevych-philip/mono-app/pkg/utils/response"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

// getProcessedStatement performs high-level validations and calls mono service to process a user's statement
func (h *Handler) getProcessedStatement(c *gin.Context) {
	var (
		token   = c.GetHeader(types.HeaderKeyXToken)
		account = c.Param(types.ParamKeyAccount)
		from    = c.Param(types.ParamKeyFrom)

		RFC3339noTimeLayout = time.RFC3339[:strings.Index(time.RFC3339, "T")]
	)

	if len(token) == 0 {
		err := fmt.Errorf("%q header is missing", types.HeaderKeyXToken)
		logrus.Error(err)
		h.utilities.ResponseHandler.CommonResponseJSON(c, http.StatusBadRequest, response.ErrorResponseKeyName, err.Error())

		return
	}
	if len(account) == 0 {
		account = "0" // default account is 0
	}
	if len(from) == 0 {
		err := fmt.Errorf("%q parameter cannot be empty", types.ParamKeyFrom)
		logrus.Error(err)
		h.utilities.ResponseHandler.CommonResponseJSON(c, http.StatusBadRequest, response.ErrorResponseKeyName, err.Error())

		return
	}

	startDate, err := time.Parse(RFC3339noTimeLayout, from)
	if err != nil {
		logrus.Error(err)
		err := fmt.Errorf("failed to parse start date: %q. Expected format is %q", from, RFC3339noTimeLayout)
		h.utilities.ResponseHandler.CommonResponseJSON(c, http.StatusBadRequest, response.ErrorResponseKeyName, err.Error())

		return
	}

	statement, err := h.services.Mono.ProcessStatement(token, account, startDate)
	if err != nil {
		h.utilities.ResponseHandler.CommonResponseJSON(c, http.StatusBadRequest, response.ErrorResponseKeyName, err.Error())

		return
	}

	h.utilities.ResponseHandler.CommonResponseJSON(c, http.StatusOK, "statement", statement)
}

func (h *Handler) getClientInfo(c *gin.Context) {
}
