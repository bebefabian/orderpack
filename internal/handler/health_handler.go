package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler handles health check requests
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
