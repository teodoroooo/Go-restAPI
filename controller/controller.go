package controller

import (
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{ID: "1", Title: "Poesia Acustica", Artist: "Orochi", Price: 356.99},
	{ID: "2", Title: "30PRAUM", Artist: "MATUE", Price: 156.99},
	{ID: "3", Title: "Mamonas", Artist: "Japinha da Nerf", Price: 256.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
