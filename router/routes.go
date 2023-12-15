package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/tiburciohugo/simple-rest-api/docs"
	"github.com/tiburciohugo/simple-rest-api/handler"
)

func initializeRoutes(r *gin.Engine) {
	basePath := "/api/v1"
	v1 := r.Group(basePath)
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		v1.GET("/health", handler.CheckHealth)
		v1.GET("/hello/*name", handler.HelloUser)
		v1.GET("/gopher/*gopher", handler.GetGopherByName)

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.Title = "Simple API"
	docs.SwaggerInfo.Description = "This is a sample API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

}
