package main

import (
	"log"
	"musical-catalog/internal/config"
	"musical-catalog/internal/models"
	"musical-catalog/internal/repository"
	"musical-catalog/internal/services"
	"musical-catalog/internal/transport"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.SetupDatabase()

	if err := db.AutoMigrate(
		&models.Artist{},
		&models.Album{},
		&models.User{},
		&models.Playlist{},
		&models.Track{},
		&models.Review{},
	); err != nil {
		log.Fatalf("auto migrate failed:  %v", err)
	}

	artistRepo := repository.NewArtistRepository(db)
	albumRepo := repository.NewAlbumRepository(db)
	userRepo := repository.NewUserRepository(db)
	playlistRepo := repository.NewPlaylistRepository(db)
	trackRepo := repository.NewTrackRepository(db)
	reviewRepo := repository.NewReviewRepository(db)

	artistService := services.NewArtistService(artistRepo)
	albumService := services.NewAlbumService(artistRepo, albumRepo)
	userService := services.NewUserService(userRepo)
	playlistService := services.NewPlaylistService(userRepo, playlistRepo, trackRepo)
	trackService := services.NewTrackService(trackRepo)
	reviewService := services.NewReviewService(reviewRepo)

	artistTransport := transport.NewArtistTransport(artistService)
	albumTransport := transport.NewAlbumTransport(albumService)
	userTransport := transport.NewUserTransport(userService)
	playlistTransport := transport.NewPlaylistTransport(playlistService)
	trackTransport := transport.NewTrackTransport(trackService)
	reviewTransport := transport.NewReviewTransport(reviewService)

	router := gin.Default()
	
	transport.RegisterRoutes(router, artistTransport, albumTransport, userTransport, playlistTransport, trackTransport, reviewTransport)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
