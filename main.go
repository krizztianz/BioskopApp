// @title Bioskop API
// @version 1.0
// @description API untuk mengelola data bioskop
// @host localhost:8080
// @BasePath /
package main

import (
	"bioskop-app/database"
	"bioskop-app/handlers"

	_ "bioskop-app/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database.InitDB()

	router := gin.Default()

	// Custom url untuk Swagger
	router.GET("/endpoints", func(c *gin.Context) {
		c.Redirect(302, "/endpoints/index.html")
	})

	// Swagger UI
	router.GET("/endpoints/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/bioskop", handlers.CreateBioskop)
	router.GET("/bioskop", handlers.GetAllBioskop)
	router.GET("/bioskop/:id", handlers.GetBioskopByID)
	router.PUT("/bioskop/:id", handlers.UpdateBioskop)
	router.DELETE("/bioskop/:id", handlers.DeleteBioskop)

	router.Run(":8080")
}
