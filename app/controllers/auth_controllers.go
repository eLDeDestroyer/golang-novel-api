package controllers

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	LoginController(ctx *fiber.Ctx) error
	RegisterController(ctx *fiber.Ctx) error
}
