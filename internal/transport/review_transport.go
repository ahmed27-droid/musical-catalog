package transport

import (
	"musical-catalog/internal/models"
	"musical-catalog/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewTransport struct {
	service services.ReviewService
}

func NewReviewTransport(service services.ReviewService) *ReviewTransport {
	return &ReviewTransport{service: service}
}

func (h *ReviewTransport) GetAll(c *gin.Context) {
	var trackID *uint

	if param := c.Query("track_id"); param != "" {
		id, err := strconv.Atoi(param)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid track_id"})
			return
		}
		t := uint(id)
		trackID = &t
	}

	reviews, err := h.service.GetAll(trackID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func (h *ReviewTransport) Create(c *gin.Context) {
	var req models.CreateReviewRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "review created"})
}

func (h *ReviewTransport) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "review deleted"})
}
