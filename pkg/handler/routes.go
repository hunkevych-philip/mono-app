package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api.monobank.ua")
	{
		personal := api.Group("/personal")
		{
			personal.GET("/client-info", h.getClientInfo)

			statement := personal.Group("/statement")
			{
				statement.GET("/:account/:from/:to", h.getClientStatement)
			}
		}
	}

	return router
}
