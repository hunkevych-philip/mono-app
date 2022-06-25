package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/hunkevych-philip/mono-app/pkg/utils/response"
)

type ResponseHandler interface {
	CommonResponseJSON(c *gin.Context, statusCode int, key string, val interface{})
}

type UtilsImpl struct {
	ResponseHandler ResponseHandler
}

func NewUtils() *UtilsImpl {
	return &UtilsImpl{
		ResponseHandler: response.NewResponseHandler(),
	}
}
