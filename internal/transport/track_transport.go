package transport

import (
	"musical-catalog/internal/models"
	"musical-catalog/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TrackTransport struct {
	service services.TrackService
}

func NewTrackTransport(service services.TrackService) *TrackTransport {
	return &TrackTransport{service: service}
}

func (h *TrackTransport) GetAll(c *gin.Context) {
	tracks, err := h.service.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tracks)
}

func (h *TrackTransport) Create(c *gin.Context) {
	var req models.CreateTrackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	track, err := h.service.Create(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, track)
}

func (h *TrackTransport) GetByID(c *gin.Context) {
	id, err:= strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(400, gin.H{"error":"invalid id"})
		return
	}

	track, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	c.JSON(200, track)
}

func (h *TrackTransport) Update(c *gin.Context) {
	id, err:= strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(400, gin.H{"error":"invalid id"})
		return
	}

	var req models.UpdateTrackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	track, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, track)
}

func (h *TrackTransport) Delete(c *gin.Context) {
	id, err:= strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(400, gin.H{"error":"invalid id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (h *TrackTransport) GetAverage(c *gin.Context) {
	id, err:= strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(400, gin.H{"error":"invalid id"})
		return
	}

	avg, err := h.service.GetAverage(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"track_id": id,
		"average":  avg,
	})
}

