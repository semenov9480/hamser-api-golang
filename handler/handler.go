package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/semenov9480/hamser-api-golang/service"
)

type Handler struct {
	service *service.Service
}

func InitHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		graph := api.Group("/graph")
		{
			graph.GET("/candels:any", h.getGraph)
		}
	}

	return router
}
