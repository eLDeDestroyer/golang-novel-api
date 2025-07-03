package main

import (
	"e-novel/config"
	"e-novel/db/seeders"
	"e-novel/router"
	"flag"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	authController, userController, bookController, pageController := config.DepedencyInjection(db)

	app := fiber.New()
	app.Static("/uploads", "./uploads")
	
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS, PATCH",
	}))

	router.SetUpRoutes(app, authController, userController, bookController, pageController)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
