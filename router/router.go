package router

import (
	"e-novel/app/controllers"
	"e-novel/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App, authController controllers.AuthController, userController controllers.UserController, bookController controllers.BookController) {
	api := app.Group("/api")
	auth := app.Group("/api/auth", middleware.AuthMiddleware())

	api.Post("/signin", authController.LoginController)
	api.Post("/signup", authController.RegisterController)

	auth.Get("/user", userController.GetMeController)
	auth.Get("/user/:username", userController.GetUserController)

	auth.Get("/book/new", bookController.GetRecentBook)
	auth.Get("/book/like", bookController.GetBookMostLike)


}
