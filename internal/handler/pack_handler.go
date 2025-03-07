package handler

import (
	"net/http"
	"strconv"

	"github.com/bebefabian/orderpack/internal/service"
	"github.com/gin-gonic/gin"
)

// PackHandler handles HTTP requests related to packs
type PackHandler struct {
	packService service.PackService
}

// NewPackHandler initializes a new PackHandler with the given service
func NewPackHandler(service service.PackService) *PackHandler {
	return &PackHandler{packService: service}
}

// GetPackSizes handles GET /packs
func (h *PackHandler) GetPackSizes(c *gin.Context) {
	packs, err := h.packService.GetPackSizes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting pack sizes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"packs": packs})
}

// UpdatePackSizes handles POST /packs
func (h *PackHandler) UpdatePackSizes(c *gin.Context) {
	var newPackSizes []int
	if err := c.BindJSON(&newPackSizes); err != nil || len(newPackSizes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pack sizes"})
		return
	}

	err := h.packService.UpdatePackSizes(newPackSizes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error updating pack sizes"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pack sizes updated", "packs": newPackSizes})

}

// CalculateOrder handles GET /calculate?quantity=X
func (h *PackHandler) CalculateOrder(c *gin.Context) {
	quantityStr := c.Query("quantity")
	orderQuantity, err := strconv.Atoi(quantityStr)
	if err != nil || orderQuantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order quantity"})
		return
	}

	// Call service method
	packDistribution, err := h.packService.CalculatePacks(orderQuantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error calculating packs"})
		return
	}

	c.JSON(http.StatusOK, packDistribution)
}
