package config

import (
	"e-novel/app/controllers"
	"e-novel/app/repositories"
	"e-novel/app/services"

	"gorm.io/gorm"
)

func DepedencyInjection(db *gorm.DB) (*controllers.AuthControllerImpl, *controllers.UserControllerImpl, *controllers.BookControllerImpl) {
	authRepository := repositories.NewAuthRepositoryImpl(db)
	authService := services.NewAuthServiceImpl(authRepository)
	authController := controllers.NewAuthControllerImpl(authService)

	userRepository := repositories.NewUserRepositoryImpl(db)
	userService := services.NewUserServiceImpl(userRepository)
	userController := controllers.NewUserControllerImpl(userService)

	bookRepository := repositories.NewBookRepositoryImpl(db)
	bookService := services.NewBookServiceImpl(bookRepository)
	bookController := controllers.NewBookControllerImpl(bookService)

	return authController, userController, bookController
}
