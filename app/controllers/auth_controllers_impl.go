package controllers

import (
	"e-novel/app/services"
	"e-novel/helper"
	"e-novel/model"
	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	authService services.AuthService
}


func NewAuthControllerImpl(authService services.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		authService: authService,
	}
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

	data, err := controller.authService.LoginService(Login.Username, Login.Password)
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

	err = controller.authService.RegisterService(&user)
	if err != nil {
		return helper.ErrorResponse(ctx, err, "fails register service")
	}

	return helper.SuccessResponse(ctx, "", "success")
}
