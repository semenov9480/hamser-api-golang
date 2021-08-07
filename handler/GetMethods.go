package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/semenov9480/hamser-api-golang/structs"
)

func (h *Handler) getCandels(c *gin.Context) {
	var input structs.GetGraphParams
	if err := c.BindJSON(&input); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Bad input params")
		return
	}
	result, err := h.service.GetCandels(input)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
		//log.Fatal(err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": result,
	})
	log.Print(c.Request.RequestURI)

}

type errorResponse struct {
	Message string `json:"message"`
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
