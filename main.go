package main

import (
	"e-novel/config"
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

	app := fiber.New()
	fmt.Println("success")

	err = app.Listen(":3001")
	if err != nil {
		panic(err)
	}
}
