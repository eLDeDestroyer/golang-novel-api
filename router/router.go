package router

import (
	"e-novel/app/controllers"
	"e-novel/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App, authController controllers.AuthController, userController controllers.UserController, pageController controllers.PageController) {
	api := app.Group("/api")
	auth := app.Group("/api/auth", middleware.AuthMiddleware())

	api.Post("/signin", authController.LoginController)
	api.Post("/signup", authController.RegisterController)

	auth.Get("/user", userController.GetMeController)
	auth.Get("/user/:username", userController.GetUserController)
	auth.Get("/user/action/:action", userController.GetBookByActionUser)
	auth.Post("/user/action/:action/:book_id", userController.AddBookActionUser)

	auth.Get("/book/page/:book_id/:page", pageController.GetPageBook)
	auth.Post("/book/page/add", pageController.AddPageBook)
	auth.Patch("/book/page/update", pageController.UpdatePageBook)

}
