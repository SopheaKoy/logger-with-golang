package v1

import (
	schema "logger/ent/schema"
	"logger/services"

	"github.com/gofiber/fiber/v2"
)

func UserHandler(app *fiber.App) (*schema.IResponseBase, error) {
    app.Post("/list", func(c *fiber.Ctx) error {
        users, err := services.ListUser()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to get users",
            })
        }
        return c.JSON(users)
    })
    
    return schema.NewIResponseBase("200", "User routes registered successfully", nil), nil
}