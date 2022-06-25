package utils

import (
	"github.com/HunkevychPhilip/todo/pkg/utils/response"
	"github.com/gin-gonic/gin"
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
