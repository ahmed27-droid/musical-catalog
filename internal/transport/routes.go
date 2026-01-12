package transport

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.Engine,
	artistTransport *ArtistTransport,
	albumTransport *AlbumTransport,
	userTransport *UserTransport,
	playlistTransport *PlaylistTransport,
	trackTransport *TrackTransport,
	reviewTransport *ReviewTransport,
	
){
	router.GET("/artists",artistTransport.ListArtists)
	router.POST("/artists", artistTransport.Create)

	router.GET("/artists/:id", artistTransport.GetArtistByID)
	router.PATCH("/artists/:id", artistTransport.UpdateArtist)
	router.DELETE("/artists/:id", artistTransport.DeleteArtist)

	router.GET("/albums", albumTransport.ListAlbums)
	router.POST("/albums", albumTransport.CreateAlbum)

	router.GET("/albums/:id", albumTransport.GetAlbumByID)
	router.PATCH("/albums/:id", albumTransport.UpdateAlbum)
	router.DELETE("/albums/:id", albumTransport.DeleteAlbum)
	router.GET("/albums/:id/average", albumTransport.GetAlbumsAverageRating)

	router.GET("/users", userTransport.List)
	router.POST("/users", userTransport.Create)

	router.GET("/users/:id", userTransport.GetByID)
	router.PATCH("/users/:id", userTransport.Update)
	router.DELETE("/user/:id", userTransport.Delete)

	router.GET("/playlists", playlistTransport.List)
	router.POST("/playlists", playlistTransport.Create)

	router.GET("/playlists/:id", playlistTransport.GetByID)
	router.DELETE("/playlists/:id", playlistTransport.Delete)

	router.POST("/playlists/:id/tracks/:track_id", playlistTransport.AddTrack)
	router.DELETE("/playlists/:id/tracks/:track_id", playlistTransport.DeleteTrack)

	router.GET("/tracks", trackTransport.GetAll)
	router.POST("/tracks", trackTransport.Create)

	router.GET("/tracks/:id", trackTransport.GetByID)
	router.PATCH("/tracks/:id", trackTransport.Update)
	router.DELETE("/tracks/:id", trackTransport.Delete)
	router.GET("/tracks/:id/average", trackTransport.GetAverage)

	router.GET("/reviews", reviewTransport.GetAll)
	router.POST("/reviews", reviewTransport.Create)
	router.DELETE("/reviews/:id", reviewTransport.Delete)


}