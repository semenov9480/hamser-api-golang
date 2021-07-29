package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getGraph(c *gin.Context) {
	test := h.service.GetGraph(c)
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": test,
	})
}
