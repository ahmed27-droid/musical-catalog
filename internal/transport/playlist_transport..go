package transport

import (
	"musical-catalog/internal/models"
	"musical-catalog/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlaylistTransport struct {
	service services.PlaylistService
}

func NewPlaylistTransport(service services.PlaylistService) *PlaylistTransport {
	return &PlaylistTransport{service: service}
}

func (t *PlaylistTransport) Create(c *gin.Context) {
	var req models.CreatePlaylist

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playlist, err := t.service.CreatePlaylist(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, playlist)
}

func (t *PlaylistTransport) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	playlist, err := t.service.GetByIDPlaylist(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, playlist)
}

func (t *PlaylistTransport) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := t.service.DeletePlaylist(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "playlist delete"})
}

func (t *PlaylistTransport) AddTrack(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	trackID, err := strconv.Atoi(c.Param("track_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := t.service.AddTrackToPlaylist(uint(id), uint(trackID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "track added"})
}

func (t *PlaylistTransport) DeleteTrack(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	trackID, err := strconv.Atoi(c.Param("track_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := t.service.AddTrackToPlaylist(uint(id), uint(trackID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "track delete"})
}

func (t *PlaylistTransport) List(c *gin.Context) {
	playlists, err := t.service.ListPlaylist()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, playlists)
}
