package main

import (
	"musical-catalog/internal/config"
	"musical-catalog/internal/models"
	"musical-catalog/internal/repository"
	"musical-catalog/internal/services"
	"musical-catalog/internal/transport"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.SetupDatabase()

	db.AutoMigrate(&models.Artist{}, &models.Album{})

	artistRepo := repository.NewArtistRepository(db)
	albumRepo := repository.NewAlbumRepository(db)

	artistService := services.NewArtistService(artistRepo, albumRepo)
	albumService := services.NewAlbumService(artistRepo, albumRepo)

	artistTransport := transport.NewArtistTransport(artistService)
	albumTransport := transport.NewAlbumTransport(albumService)

	router := gin.Default()
	transport.RegisterRoutes(router, artistTransport, albumTransport)
	router.Run(":8080")

}
