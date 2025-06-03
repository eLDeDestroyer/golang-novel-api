package helper

import "github.com/gofiber/fiber/v2"

func SuccessResponse(ctx *fiber.Ctx, data interface{}, message string) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
		"message": message,
	})
}

func ErrorResponse(ctx *fiber.Ctx, err error, message string) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"success": false,
		"error":   err,
		"message": message,
	})
}
