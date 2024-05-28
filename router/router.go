package router

import (
	"rest-api/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()

	router.GET("/albums", controller.GetAlbums)

	router.Run("localhost:8080")

}
