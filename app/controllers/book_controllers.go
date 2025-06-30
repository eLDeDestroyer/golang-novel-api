package controllers

import "github.com/gofiber/fiber/v2"

type BookController interface {
	GetCategories(ctx *fiber.Ctx) error

	GetBookByUsername(ctx *fiber.Ctx) error	
	GetRecentBook(ctx *fiber.Ctx) error
	GetBookMostLike(ctx *fiber.Ctx) error
	GetBookByCategoryId(ctx *fiber.Ctx) error
	GetBookDetailById(ctx *fiber.Ctx) error

	AddBook(ctx *fiber.Ctx) error
	AddBookCategory(ctx *fiber.Ctx) error

	UpdateBook(ctx *fiber.Ctx) error
	UpdateBookCategory(ctx *fiber.Ctx) error

	DeleteBook(ctx *fiber.Ctx) error

}