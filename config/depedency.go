package config

import (
	"e-novel/app/controllers"
	"e-novel/app/repositories"
	"e-novel/app/services"

	"gorm.io/gorm"
)


func DepedencyInjection(db *gorm.DB) (*controllers.AuthControllerImpl) {
	authRepository := repositories.NewAuthRepositoryImpl(db)
	authService := services.NewAuthServiceImpl(authRepository)
	authController := controllers.NewAuthControllerImpl(authService)

	return authController
}