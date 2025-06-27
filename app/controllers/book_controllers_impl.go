package controllers

import (
	"e-novel/app/services"
	"e-novel/helper"

	"github.com/gofiber/fiber/v2"
)

type BookControllerImpl struct {
	bookService services.BookService
}

func NewBookControllerImpl(bookService services.BookService) *BookControllerImpl {
	return &BookControllerImpl{
		bookService: bookService,
	}
}

func (controller *BookControllerImpl) GetRecentBook(ctx *fiber.Ctx) error {
	data, err := controller.bookService.GetRecentBook()
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get data")
	}

	return helper.SuccessResponse(ctx, data, "Success get data")
}

func (controller *BookControllerImpl) GetBookMostLike(ctx *fiber.Ctx) error {
	data, err := controller.bookService.GetBookMostLike()
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get data")
	}

	return helper.SuccessResponse(ctx, data, "Success get data")
}