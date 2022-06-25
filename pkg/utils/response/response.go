package response

import (
	"github.com/gin-gonic/gin"
)

type ResponseHandlerImpl struct {
}

func NewResponseHandler() *ResponseHandlerImpl {
	return &ResponseHandlerImpl{}
}

func (r *ResponseHandlerImpl) CommonResponseJSON(c *gin.Context, statusCode int, key string, val interface{}) {

}
