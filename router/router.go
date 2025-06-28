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

	auth.Get("/categories", bookController.GetCategories)

	auth.Get("/book/new", bookController.GetRecentBook)
	auth.Get("/book/like", bookController.GetBookMostLike)
	auth.Get("/book/:id", bookController.GetBookByCategoryId)

	auth.Get("/book/detail/:id", bookController.GetBookDetailById)

	auth.Post("/book/add", bookController.AddBook)
	auth.Post("/book/categories/add", bookController.AddBookCategory)

	auth.Patch("/book/update/:book_id", bookController.UpdateBook)
	auth.Post("/book/categories/update", bookController.UpdateBookCategory)

	auth.Delete("/book/delete/:book_id", bookController.DeleteBook)


}
