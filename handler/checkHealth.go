package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckHealth godoc
// @Summary Check health of the API
// @Description Check if the API is healthy
// @ID check-health
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "healthy",
	})
}
