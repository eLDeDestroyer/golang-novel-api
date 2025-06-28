package controllers

import "github.com/gofiber/fiber/v2"

type UserController interface {
	GetMeController(ctx *fiber.Ctx) error
	GetUserController(ctx *fiber.Ctx) error
	GetBookByActionUser(ctx *fiber.Ctx) error
	AddBookActionUser(ctx *fiber.Ctx) error

}