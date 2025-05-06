package auth

import "github.com/gofiber/fiber/v2"

func TokenAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		// Check for Bearer prefix
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or malformed Authorization header",
			})
		}

		token := authHeader[7:] // Remove "Bearer " prefix

		if token != "12324" { // Replace with your real token
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Continue if token is valid
		return c.Next()
	}
}
