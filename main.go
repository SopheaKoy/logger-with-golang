package main

import (
	"fmt"
	"logger/config"
	_ "logger/docs"
	"logger/middlewares"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"logger/handler/auth"
)

// #========================
// CreatePublicRequest represents the request body for creating a public resource
type CreatePublicRequest struct {
    Name        string `json:"name" example:"Resource name"`
    Description string `json:"description" example:"Resource description"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
    Message string `json:"message" example:"Error message"`
    Code    int    `json:"code" example:"400"`
}

// @title My API
// @version 1.0
// @description This is a simple API to demonstrate Swagger integration in Go with Fiber.
// @host localhost:3000
// @BasePath /api/v1
func main() {

	// Create a new logger instance from your custom logger
	log := config.NewLogger()

	// Create an instance of Fiber
	app := fiber.New(fiber.Config{
		StrictRouting : true,
		ErrorHandler  : middlewares.HTTPExceptionHandler,
	})

	// #=================== swagger configuration
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Configure Swagger
	app.Get("/swagger/*", swagger.New(swagger.Config{
		Title		: "Logger Service",		
		URL			: "/swagger/doc.json",
		DeepLinking	: true,
	}))
	
	// Apply the logging middleware
	// Create log middleware instance with the logger
	logMiddleware := middlewares.NewLogMiddleware(log)
	app.Use(logMiddleware.LogAccess())

	// Add group API with /api/v1 prefix
	apiPrefix := app.Group("/api/v1")

	apiPrefix.Get("/public", publicHandler)
	apiPrefix.Post("/public", publicCreationHandler)

	apiPrefix.Get("/user", userHandler)

	// Upload Service
	apiPrefix.Post("/upload", auth.TokenAuthMiddleware(), fileHandler)

	// Log the starting message
	log.Info("Starting server on port 3000...")

	// Start the app
	if err := app.Listen(":3000"); err != nil {
		log.Error().Error("Failed to start server:", err)
	}
}


// @Summary Public route
// @Description This is a public route that is accessible by everyone
// @Tags Public
// @Produce text/plain
// @Success 200 {string} string "Hello, Public Page....!!!"
// @Router /public [get]
func publicHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, Public Page....!!!")
}

// @Summary Create public resource
// @Description Creates a new public resource in the system
// @Tags Public
// @Accept json
// @Produce text/plain
// @Param data body CreatePublicRequest true "Resource creation data"
// @Success 201 {string} string "Resource created successfully"
// @Failure 400 {object} ErrorResponse "Invalid request data"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /public [post]
func publicCreationHandler(c *fiber.Ctx) error {
    return c.Status(fiber.StatusCreated).SendString("Hello, Public Page....!!!")
}

// @Tags User
// @Router /user [get]
func userHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, Public Page....!!!")
}

// fileHandler handles file uploads.
// @Summary Upload a file
// @Description Uploads a file to the server
// @Tags File Upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {string} string "Upload successful"
// @Failure 400 {string} string "Failed to upload file"
// @Security ApiKeyAuth
// @Router /upload [post]
func fileHandler(c *fiber.Ctx) error {
	// Get the uploaded file from the form
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "File not provided",
		})
	}

	// Optional: Create the upload directory if it doesn't exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create upload directory",
		})
	}

	// Save the file
	savePath := fmt.Sprintf("./uploads/%s", fileHeader.Filename)
	err = c.SaveFile(fileHeader, savePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}

	// Success response
	return c.JSON(fiber.Map{
		"message": "Upload successful",
		"file":    fileHeader.Filename,
	})
}
