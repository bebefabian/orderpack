package main

import "github.com/bebefabian/orderpack/cmd/app"

func main() {
	// Create and initialize the app
	application := &app.App{}
	application.Initialize()

	// Start the server on port 8080
	application.Run("8080")
}
