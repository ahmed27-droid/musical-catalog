package transport

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.Engine,
	artistTransport *ArtistTransport,
	albumTransport *AlbumTransport,
	
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

}