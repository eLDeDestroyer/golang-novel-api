package router

import (
	"e-novel/app/controllers"
	// "e-novel/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App, authController controllers.AuthController)  {
	api := app.Group("/api")
	// auth := app.Group("/api/auth", middleware.AuthMiddleware())

	api.Post("/signin", authController.LoginController)
	api.Post("/signup", authController.RegisterController)
}