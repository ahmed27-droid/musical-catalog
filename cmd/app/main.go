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

	db.AutoMigrate(&models.Artist{}, &models.Album{}, &models.User{}, &models.Playlist{})

	artistRepo := repository.NewArtistRepository(db)
	albumRepo := repository.NewAlbumRepository(db)
	userRepo := repository.NewUserRepository(db)
	playlistRepo := repository.NewPlaylistRepository(db)

	artistService := services.NewArtistService(artistRepo)
	albumService := services.NewAlbumService(artistRepo, albumRepo)
	userService := services.NewUserService(userRepo)
	playlistService := services.NewPlaylistService(userRepo, playlistRepo, trackRepo)

	artistTransport := transport.NewArtistTransport(artistService)
	albumTransport := transport.NewAlbumTransport(albumService)
	userTransport := transport.NewUserTransport(userService)
	playlistTransport :=transport.NewPlaylistTransport(playlistService)


	router := gin.Default()
	transport.RegisterRoutes(router, artistTransport, albumTransport, userTransport, playlistTransport)
	router.Run(":8080")

}
