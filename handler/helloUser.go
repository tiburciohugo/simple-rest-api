package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// MessageResponse defines the response structure
// @name MessageResponse
// @schema
//
//	{
//	  "type": "object",
//	  "properties": {
//	    "message": {
//	      "type": "string"
//	    }
//	  }
//	}
type MessageResponse struct {
	Message string `json:"message"`
}

// HelloUser godoc
// @Summary Greet a user
// @Description greet a user by name
// @ID hello-user
// @Accept  json
// @Produce  json
// @Param   name     path    string     true        "User name"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} MessageResponse
// @Router /hello/{name} [get]
func HelloUser(c *gin.Context) {
	name := strings.TrimPrefix(c.Param("name"), "/")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "name is required",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, " + name + "!",
	})
}
