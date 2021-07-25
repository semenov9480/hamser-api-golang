package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		graph := api.Group("/graph")
		{
			graph.GET("/candels", h.getCandels)
		}
	}

	return router
}
