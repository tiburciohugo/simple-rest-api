package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetGopherByName godoc
// @Summary Get gopher image by name
// @Description Redirects to the URL of a gopher image based on the provided gopher name
// @ID get-gopher-by-name
// @Accept  json
// @Produce  json
// @Param   gopher     path    string     true        "Gopher name"
// @Success 301 {string} string "Redirects to the URL of the gopher image"
// @Router /gopher/{gopher} [get]
func GetGopherByName(c *gin.Context) {
	gopher := strings.TrimPrefix(c.Param("gopher"), "/")
	var URL string
	if gopher != "" {
		URL = "https://github.com/scraly/gophers/raw/main/" + gopher + ".png"
	} else {
		URL = "https://github.com/scraly/gophers/raw/main/dr-who.png"
	}
	c.Redirect(http.StatusMovedPermanently, URL)
}
