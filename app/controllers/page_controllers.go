package controllers

import "github.com/gofiber/fiber/v2"

type PageController interface {
	GetPageBook(ctx *fiber.Ctx) error
	AddPageBook(ctx *fiber.Ctx) error
	UpdatePageBook(ctx *fiber.Ctx) error

}