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
			// get
			graph.GET("/candels", h.getCandels)
			//graph.GET("/candelPatterns", h.GetCandels)
			//graph.GET("/trendLines", h.GetCandels)
			//graph.GET("/trendLevels", h.GetCandels)
			//graph.GET("/trend", h.GetCandels)
			// post
			//graph.POST("/klines", h.GetCandels)
			///graph.POST("/candelPatterns", h.GetCandels)
			//graph.POST("/trendLines", h.GetCandels)
			//graph.POST("/trendLevels", h.GetCandels)
			//graph.POST("/trend", h.GetCandels)
		}
		//user := api.Group("/user")
		//{
		// access by tokens
		//user.GET("/jobs", h.GetCandels)
		//user.GET("/jobInfo", h.GetCandels)
		//}
	}

	return router
}
