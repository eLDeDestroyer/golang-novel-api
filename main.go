package main

import (
	"e-novel/config"
	"e-novel/router"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()
	db := config.ConnectDB()

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	authController := config.DepedencyInjection(db)

	app := fiber.New()
	fmt.Println("success")

	router.SetUpRoutes(app,authController)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
