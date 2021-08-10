package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/semenov9480/hamser-api-golang/structs"
)

func (h *Handler) getCandels(c *gin.Context) {
	var input structs.GetGraphParams
	if err := c.BindQuery(&input); err != nil {

		ErrorResponse(c, http.StatusBadRequest, "Bad input params")
		return
	}
	fmt.Print(input)
	result, err := h.service.GetCandels(input)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
		//log.Fatal(err)
	}
	c.JSON(http.StatusOK, result)
	log.Print(c.Request.RequestURI)
}

func (h *Handler) getLines(c *gin.Context) {
	var input structs.GetToolsParams
	if err := c.BindQuery(&input); err != nil {

		ErrorResponse(c, http.StatusBadRequest, "Bad input params")
		return
	}
	fmt.Print(input)

	result, err := h.service.GetLines(input)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
		//log.Fatal(err)
	}
	c.JSON(http.StatusOK, result)
	log.Print(c.Request.RequestURI)
}

func (h *Handler) getLevels(c *gin.Context) {
	var input structs.GetToolsParams
	if err := c.BindQuery(&input); err != nil {

		ErrorResponse(c, http.StatusBadRequest, "Bad input params")
		return
	}
	fmt.Print(input)

	result, err := h.service.GetLevels(input)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
		//log.Fatal(err)
	}
	c.JSON(http.StatusOK, result)
	log.Print(c.Request.RequestURI)
}

func (h *Handler) getPatterns(c *gin.Context) {
	var input structs.GetToolsParams
	if err := c.BindQuery(&input); err != nil {

		ErrorResponse(c, http.StatusBadRequest, "Bad input params")
		return
	}
	fmt.Print(input)

	result, err := h.service.GetPatterns(input)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
		//log.Fatal(err)
	}
	c.JSON(http.StatusOK, result)
	log.Print(c.Request.RequestURI)
}

type errorResponse struct {
	Message string `json:"message"`
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
