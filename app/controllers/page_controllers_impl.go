package controllers

import (
	"e-novel/app/services"
	"e-novel/helper"
	"e-novel/model"

	"github.com/gofiber/fiber/v2"
)

type PageControllerImpl struct {
	pageService services.PageService
}

func NewPageControllerImpl(pageService services.PageService) *PageControllerImpl {
	return &PageControllerImpl{
		pageService: pageService,
	}
}


func (controller *PageControllerImpl) GetPageBook(ctx *fiber.Ctx) error {
	bookId,err := ctx.ParamsInt("book_id")
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get book id")
	}

	page,err := ctx.ParamsInt("page")
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get page params")
	}

	data , err := controller.pageService.GetPageBook(bookId, page)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get page book")
	}

	return helper.SuccessResponse(ctx, data, "success get pages")
}


func (controller *PageControllerImpl) AddPageBook(ctx *fiber.Ctx) error {
	var data model.Page

	err := ctx.BodyParser(&data)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails  parse")
	}

	err = controller.pageService.AddPageBook(&data)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails add page")
	}

	return helper.SuccessResponse(ctx, data, "success add page")
}


func (controller *PageControllerImpl) UpdatePageBook(ctx *fiber.Ctx) error {
	var data model.Page

	err := ctx.BodyParser(&data)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails  parse")
	}

	err = controller.pageService.UpdatePageBook(&data)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails update page")
	}

	return helper.SuccessResponse(ctx, data, "success update page")
}

