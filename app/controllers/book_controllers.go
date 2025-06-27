package controllers

import "github.com/gofiber/fiber/v2"

type BookController interface {
	GetRecentBook(ctx *fiber.Ctx) error
	GetBookMostLike(ctx *fiber.Ctx) error

}