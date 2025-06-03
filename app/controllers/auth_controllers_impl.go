package controllers

import (
	"e-novel/app/services"
	"e-novel/helper"
	"e-novel/model"
	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	authController services.AuthService
}

func (controller *AuthControllerImpl) LoginController(ctx *fiber.Ctx) error {
	var Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := ctx.BodyParser(&Login)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails parser")
	}

	data, err := controller.authController.LoginService(Login.Username, Login.Password)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails login")
	}

	return helper.SuccessResponse(ctx, data, "success")
}

func (controller *AuthControllerImpl) RegisterController(ctx *fiber.Ctx) error {
	var user model.User

	err := ctx.BodyParser(&user)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails body parser")
	}

	err = controller.authController.RegisterService(&user)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails register service")
	}

	return helper.SuccessResponse(ctx, token, "success")
}
