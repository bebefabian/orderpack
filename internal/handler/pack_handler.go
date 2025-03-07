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
	packs := h.packService.GetPackSizes()
	c.JSON(http.StatusOK, gin.H{"packs": packs})
}

// UpdatePackSizes handles POST /packs
func (h *PackHandler) UpdatePackSizes(c *gin.Context) {
	var newPackSizes []int
	if err := c.BindJSON(&newPackSizes); err != nil || len(newPackSizes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pack sizes"})
		return
	}

	h.packService.UpdatePackSizes(newPackSizes)
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
	packDistribution := h.packService.CalculatePacks(orderQuantity)

	c.JSON(http.StatusOK, packDistribution)
}
