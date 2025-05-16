package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"logger/config"

	schema "logger/ent/schema"
)

// LogMiddleware struct that holds the Logger and MaintenanceMode
type LogMiddleware struct {
	Logger          *config.Logger // Logger is now from the config package
	MaintenanceMode bool
}

// NewLogMiddleware creates a new instance of LogMiddleware with the logger passed in
func NewLogMiddleware(logger *config.Logger) *LogMiddleware {
	return &LogMiddleware{
		Logger:          logger,      // Pass logger to LogMiddleware
		MaintenanceMode: false,       // Default to false; modify as needed
	}
}

func (lm *LogMiddleware) LogAccess() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if lm.MaintenanceMode {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"message": "Server is currently under maintenance. Please try again later.",
			})
		}

		contentType := c.Get("Content-Type")
		method      := c.Method()
		url         := c.OriginalURL()
		var body string
	
		if method == fiber.MethodPost || method == fiber.MethodPut {
			switch {
			case strings.HasPrefix(contentType, fiber.MIMEApplicationJSON):
				body = string(c.Body())

			case strings.HasPrefix(contentType, fiber.MIMEApplicationForm):
				body = string(c.Body())

			case strings.HasPrefix(contentType, fiber.MIMEMultipartForm):
				form, err := c.MultipartForm()
				if err == nil && form != nil {
					formData := make(map[string][]string)
					for key, val := range form.Value {
						formData[key] = val
					}
					jsonBytes, _ := json.Marshal(formData)
					body = string(jsonBytes)
				} else {
					lm.Logger.Error().Error("Error parsing multipart form")
				}
			}
		}
		
		start        := time.Now()
		err          := c.Next()
		durationInMs := float64((time.Since(start)).Nanoseconds()) / 1e6
		statusCode   := c.Response().StatusCode()

		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		
		logEntry := fmt.Sprintf("%s %s - %d - %.2f ms | Body: %s",
			method, url, statusCode, durationInMs, body,
		)

		fmt.Println("LOGGER FROM MIDDLEWARE=", logEntry)

		if c.Method() != fiber.MethodOptions {
			if err != nil || statusCode >= 400 {
				lm.Logger.Error().Error(logEntry)
			} else {
				lm.Logger.Info(logEntry)
			}
		}
		return err
	}
}

// HTTPExceptionHandler handles custom error responses globally
func HTTPExceptionHandler(c *fiber.Ctx, err error) error {
	
	// Check if the error is of type *fiber.Error (HTTP-related error)
	if fiberErr, ok := err.(*fiber.Error); ok {
		// Log the error details (could also log it to a file or external system)
		log.Printf("HTTP Error: %v, StatusCode: %d", err.Error(), fiberErr.Code)

		// Return a custom JSON response for HTTP exceptions
		return c.Status(fiberErr.Code).JSON(schema.IResponseBase{
			LogID	: uuid.New(),
			Success	: 0, 
			Code	: fmt.Sprintf("%d", fiberErr.Code),
			Message	: fiberErr.Message,
		})
	}

	// Handle non-HTTP errors by returning a 500 Internal Server Error
	log.Printf("Unknown Error: %v", err)

	// Return a generic 500 Internal Server Error response with the error details
	return c.Status(fiber.StatusInternalServerError).JSON(schema.IResponseBase{
		LogID	: uuid.New(),
		Success	: 0,
		Code	: "500",
		Message	: "Internal Server Error",
		Data	: fmt.Sprintf("%v", err),
	})
}

// RequestValidationErrorHandler handles validator.ValidationErrors
func RequestValidationErrorHandler(c *fiber.Ctx, err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errs := make(map[string]string)
		for _, fieldErr := range validationErrors {
			errs[fieldErr.Field()] = fmt.Sprintf("must be %s", fieldErr.Tag())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validation_errors": errs,
		})
	}

	return HTTPExceptionHandler(c, err)
}