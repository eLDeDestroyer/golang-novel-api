package controllers

import (
	"e-novel/app/services"
	"e-novel/helper"
	"e-novel/model/dto"

	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	userService services.UserService
}

func NewUserControllerImpl(userService services.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (controller *UserControllerImpl) GetMeController(ctx *fiber.Ctx) error {
	userId := int(ctx.Locals("id").(uint))

	data, err := controller.userService.GetMeService(userId)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get data")
	}

	return helper.SuccessResponse(ctx, data, "success get Data")
}

func (controller *UserControllerImpl) GetUserController(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	data, err := controller.userService.GetUserService(username)

	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get data")
	}

	return helper.SuccessResponse(ctx, data, "success get Data User")
}

func (controller *UserControllerImpl) GetBookByActionUser(ctx *fiber.Ctx) error {
	action := ctx.Params("action")
	userId := int(ctx.Locals("id").(uint))

	data, err := controller.userService.GetBookByActionUser(action, userId)

	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get data")
	}

	return helper.SuccessResponse(ctx, data, "success get Data User")
}

func (controller *UserControllerImpl) AddBookActionUser(ctx *fiber.Ctx) error {
	userId := int(ctx.Locals("id").(uint))
	action := ctx.Params("action")

	bookId, err := ctx.ParamsInt("book_id")
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails get book_id")
	}

	data := dto.RequestAction{
		UserId: userId,
		BookId: bookId,
	}

	err = controller.userService.AddBookActionUser(action, &data)

	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails add actioon")
	}

	return helper.SuccessResponse(ctx, nil, "success add Data User")
}
