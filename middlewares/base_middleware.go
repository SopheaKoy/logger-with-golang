package middlewares

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"logger/config"
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
		
		start        := time.Now()
		err          := c.Next()
		durationInMs := float64((time.Since(start)).Nanoseconds()) / 1e6 // Convert nanoseconds to milliseconds
		statusCode   := c.Response().StatusCode()
		
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		accessLogMessage := fmt.Sprintf("%s %s - %d - %.2f ms",
			c.Method(),
			c.OriginalURL(),
			statusCode,
			durationInMs,
		)

		fmt.Println(accessLogMessage)

		if c.Method() != fiber.MethodOptions {
			if statusCode != fiber.StatusOK {
				lm.Logger.Error().Error(accessLogMessage)
			} else {
				lm.Logger.Info(accessLogMessage)
			}
		}

		return err
	}
}

// HTTPExceptionHandler handles general HTTP errors and logs them
func HTTPExceptionHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msg := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		msg = e.Message
	}

	return c.Status(code).JSON(fiber.Map{
		"error": msg,
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