package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bebefabian/orderpack/internal/handler"
	"github.com/bebefabian/orderpack/internal/repository"
	"github.com/bebefabian/orderpack/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// App struct holds router and dependencies
type App struct {
	Router *gin.Engine
}

// Initialize sets up the app with routes and dependencies
func (a *App) Initialize() {
	// Initialize repository, service, and handlers
	repo := repository.NewMemoryPackRepository()
	packService := service.NewPackService(repo)
	packHandler := handler.NewPackHandler(packService)

	// Initialize Gin router
	a.Router = gin.Default()

	// Enable CORS
	a.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://orderpack-production.up.railway.app"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// Handle OPTIONS method for CORS
	a.Router.OPTIONS("/*any", func(c *gin.Context) {
		c.Status(204)
	})

	// Register routes
	a.Router.GET("/packs", packHandler.GetPackSizes)
	a.Router.POST("/packs", packHandler.UpdatePackSizes)
	a.Router.GET("/calculate", packHandler.CalculateOrder)
	a.Router.GET("/health", handler.HealthHandler)
}

// Run starts the server
func (a *App) Run(port string) {
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        a.Router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Server running on port %s...", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
