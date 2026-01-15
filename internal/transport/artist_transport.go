package transport

import (
	"musical-catalog/internal/models"
	"musical-catalog/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArtistTransport struct {
	service services.ArtistService
}

func NewArtistTransport(service services.ArtistService) *ArtistTransport {
	return &ArtistTransport{service: service}
}

func (t *ArtistTransport) Create(c *gin.Context) {
	var req models.ArtistCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	artist, err := t.service.CreateArtist(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, artist)

}

func (t *ArtistTransport) ListArtists(c *gin.Context) {

	
	artists, err := t.service.ListArtist()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, artists)

}

func (t *ArtistTransport) GetArtistByID(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil || idInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	artist, err := t.service.GetArtist(uint(idInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, artist)
}

func (t *ArtistTransport) UpdateArtist(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil || idInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID",
		})
		return
	}

	var req models.ArtistUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	artist, err := t.service.UpdateArtist(uint(idInt), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, artist)

}

func (t *ArtistTransport) DeleteArtist(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil || idInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID",
		})
		return
	}

	if err := t.service.DeleteArtist(uint(idInt)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)

}
