package transport

import (
	"musical-catalog/internal/models"
	"musical-catalog/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

type AlbumTransport struct {
	service services.AlbumService
}

func NewAlbumTransport(service services.AlbumService) *AlbumTransport {
	return &AlbumTransport{service: service}
}

func (t *AlbumTransport) ListAlbums(c *gin.Context) {
	albums, err := t.service.ListAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func (t *AlbumTransport) CreateAlbum(c *gin.Context) {
	var req models.AlbumCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	album, err := t.service.CreateAlbum(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, album)
}

func (t *AlbumTransport) GetAlbumByID(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil || idInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	album, err := t.service.GetAlbum(uint(idInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, album)
}

func (t *AlbumTransport) GetAlbumsAverageRating(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil || idInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	avg, err := t.service.GetAlbumsAverageRating(uint(idInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"album_id":       idInt,
		"average_rating": avg,
	})

}

func (t *AlbumTransport) UpdateAlbum(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil || idInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID",
		})
		return
	}

	var req models.AlbumUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	album, err := t.service.UpdateAlbum(uint(idInt), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, album)
}

func (t *AlbumTransport) DeleteAlbum(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil || idInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID",
		})
		return
	}

	if err := t.service.DeleteAlbum(uint(idInt)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)

}
