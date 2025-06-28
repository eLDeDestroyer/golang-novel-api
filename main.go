package main

import (
	"e-novel/config"
	"e-novel/db/seeders"
	"e-novel/router"
	"flag"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	seed := flag.Bool("seed", false, "seed")
	flag.Parse()

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

	if *seed {
		seeders.Seed(db)
		fmt.Println("Seeder Buku berhasil dijalankan!")
		os.Exit(0)
	}

	authController, userController, pageController := config.DepedencyInjection(db)

	app := fiber.New()
	fmt.Println("success")

	router.SetUpRoutes(app, authController, userController, pageController)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
