package handler

import (
	"github.com/bebefabian/orderpack/internal/repository"
	"github.com/bebefabian/orderpack/internal/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Setup a test server
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode) // Run in test mode
	router := gin.Default()

	// Mock repository and service
	mockRepo := repository.NewMemoryPackRepository()
	mockService := service.NewPackService(mockRepo)
	packHandler := NewPackHandler(mockService)

	// Register routes
	router.GET("/packs", packHandler.GetPackSizes)
	router.POST("/packs", packHandler.UpdatePackSizes)
	router.GET("/calculate", packHandler.CalculateOrder)

	return router
}

// Test Get Packs (Initially Empty)
func TestGetPacks_Empty(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/packs", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.JSONEq(t, `{"packs":[]}`, resp.Body.String()) // Should return an empty list
}

// Test Update Packs with Invalid JSON
func TestUpdatePacks_InvalidInput(t *testing.T) {
	router := setupRouter()
	body := `"invalid_json"`

	req, _ := http.NewRequest("POST", "/packs", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

// Test Calculate Order
func TestCalculateOrder(t *testing.T) {
	router := setupRouter()

	// Set pack sizes first
	_, _ = http.NewRequest("POST", "/packs", strings.NewReader(`["250", "500", "1000"]`))

	req, _ := http.NewRequest("GET", "/calculate?quantity=1200", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), `"orderQuantity":1200`)
}

// Test Calculate Order with Invalid Input
func TestCalculateOrder_InvalidInput(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/calculate?quantity=abc", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
