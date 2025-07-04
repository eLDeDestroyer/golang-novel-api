package controllers

import (
	"e-novel/app/services"
	"e-novel/helper"
	"e-novel/model/dto"

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

func (controller *BookControllerImpl) GetCategories(ctx *fiber.Ctx) error {
	data, err := controller.bookService.GetCategories()
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get data")
	}

	return helper.SuccessResponse(ctx, data, "Success get data")
}

func (controller *BookControllerImpl) GetBookByUsername(ctx *fiber.Ctx) error {
	title := ctx.Query("title")

	data, err := controller.bookService.GetBookByUsername(title)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get data")
	}

	return helper.SuccessResponse(ctx, data, "Success get data")
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

func (controller *BookControllerImpl) GetBookByCategoryId(ctx *fiber.Ctx) error {
	categoryId, err := ctx.ParamsInt("id")
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get category id")
	}

	data, err := controller.bookService.GetBookByCategoryId(categoryId)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get data")
	}

	return helper.SuccessResponse(ctx, data, "Success get data")
}

func (controller *BookControllerImpl) GetBookDetailById(ctx *fiber.Ctx) error {
	bookId, err := ctx.ParamsInt("id")
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get category id")
	}

	data, err := controller.bookService.GetBookDetailById(bookId)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get data")
	}

	return helper.SuccessResponse(ctx, data, "Success get data")
}

func (controller *BookControllerImpl) AddBook(ctx *fiber.Ctx) error {
	var valueStruct dto.RequestBook
	userId := ctx.Locals("id").(uint)

	err := ctx.BodyParser(&valueStruct)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get value")
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get image")
	}

	bookId, err := controller.bookService.AddBook(ctx, &valueStruct, file, int(userId))
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails add book")
	}

	return helper.SuccessResponse(ctx, bookId, "success add data")

}

func (controller *BookControllerImpl) AddBookCategory(ctx *fiber.Ctx) error {
	var data dto.RequestBookCategory

	err := ctx.BodyParser(&data)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails parse data")
	}

	err = controller.bookService.AddBookCategory(&data)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails add data book category")
	}

	return helper.SuccessResponse(ctx, data, "success add data")
}

func (controller *BookControllerImpl) UpdateBook(ctx *fiber.Ctx) error {
	var valueStruct dto.RequestBook
	userId := ctx.Locals("id").(uint)
	bookId, err := ctx.ParamsInt("book_id")
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get book_id")
	}

	err = ctx.BodyParser(&valueStruct)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get value")
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get image")
	}

	data, err := controller.bookService.UpdateBook(ctx, &valueStruct, file, int(userId), bookId)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails update book")
	}

	return helper.SuccessResponse(ctx, data, "success update data")
}

func (controller *BookControllerImpl) UpdateBookCategory(ctx *fiber.Ctx) error {
	var data dto.RequestBookCategory

	err := ctx.BodyParser(&data)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails parse data")
	}

	err = controller.bookService.UpdateBookCategory(&data)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails add data book category")
	}

	return helper.SuccessResponse(ctx, data, "success add data")
}

func (controller *BookControllerImpl) DeleteBook(ctx *fiber.Ctx) error {
	bookId, err := ctx.ParamsInt("book_id")

	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get book_id")
	}

	err = controller.bookService.DeleteBook(bookId)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails  delete book")
	}

	return helper.SuccessResponse(ctx, nil, "success add data")
}
