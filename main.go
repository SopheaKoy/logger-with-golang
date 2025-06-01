package main

import (
	"fmt"
	_ "logger_with_go/docs"
	"logger_with_go/middlewares"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	config "logger_with_go/config"
	"logger_with_go/db"
	"logger_with_go/handler/auth"
	schema "logger_with_go/schemas"
)

type CreatePublicRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
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

	// connection db
	db.InitDB()

	// Create a new logger instance from your custom logger
	log := config.NewLogger()
	logMiddleware := middlewares.NewLogMiddleware(log)

	// Create an instance of Fiber
	app := fiber.New(fiber.Config{
		StrictRouting: true,
		ErrorHandler:  middlewares.HTTPExceptionHandler,
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	// Configure Swagger
	app.Get("/swagger/*", swagger.New(swagger.Config{
		Title:       "Logger Service",
		URL:         "/swagger/doc.json",
		DeepLinking: true,
	}))

	// CORS middleware setup (allow all origins)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.LoadSettings().CORS_ALLOWED_ORIGINS,
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "*",
		AllowCredentials: false,
	}))

	// Apply the logging middleware
	app.Use(logMiddleware.LogAccess())

	// Register error handlers
	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			if validationErr, ok := err.(validator.ValidationErrors); ok {
				return middlewares.RequestValidationErrorHandler(c, validationErr)
			}
			return middlewares.HTTPExceptionHandler(c, err)
		}
		return nil
	})

	// Add group API with /api/v1 prefix
	apiPrefix := app.Group(config.LoadSettings().API_PREFIX_V1)

	apiPrefix.Get("/public", publicHandler)
	apiPrefix.Post("/public", publicCreationHandler)
	apiPrefix.Get("/user", userHandler)
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
	// Your implementation here
	return c.Status(fiber.StatusCreated).SendString("Hello, Public Page....!!!")
}

// @Tags User
// @Router /user [get]
func userHandler(c *fiber.Ctx) error {
	// send notify

	response := schema.NewResponse()
	response.Code = "200"
	response.Data = "Getting start call ther user....!!!"

	return c.Status(fiber.StatusOK).JSON(response)
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
		response := schema.NewErrorResponse(err)
		response.Code = "400"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Optional: Create the upload directory if it doesn't exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		response := schema.NewErrorResponse(err)
		response.Code = fmt.Sprintf("%d", fiber.StatusOK)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Save the file
	savePath := fmt.Sprintf("./uploads/%s", fileHeader.Filename)
	err = c.SaveFile(fileHeader, savePath)
	if err != nil {
		response := schema.NewErrorResponse(err)
		response.Code = fmt.Sprintf("%d", fiber.StatusOK)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Success response
	response := schema.NewResponse()
	response.Code = fmt.Sprintf("%d", fiber.StatusOK)
	response.Data = fiber.Map{
		"message": "Upload successful",
		"file":    fileHeader.Filename,
	}
	return c.JSON(response)
}
