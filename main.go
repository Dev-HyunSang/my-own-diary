package main

import (
	"log"
	"strings"

	"github.com/dev-hyunsang/my-own-diary/cmd"
	"github.com/dev-hyunsang/my-own-diary/config"
	"github.com/dev-hyunsang/my-own-diary/database"
	"github.com/dev-hyunsang/my-own-diary/model"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	db, err := database.ConnectionDB()
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed DataBase Connection")
		log.Fatalln(err.Error())
	}

	err = db.AutoMigrate(&model.Users{})
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed to DataBase AutoMigrate")
		log.Fatalln(err)
	}

	app := fiber.New(fiber.Config{
		AppName: "my own diary",
	})

	app.Use(cors.New(cors.Config{
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodDelete,
		}, ","),
	}))

	app.Post("/register", cmd.RegisterHandler)
	app.Get("/login", cmd.LoginedIndexHandler)

	user := app.Group("/user")
	user.Use(jwtware.New(jwtware.Config{
		SigningKey: config.GetEnv("SIGNING_KEY"),
	}))
	user.Get("/", cmd.LoginedIndexHandler)

	if err = app.Listen(":3000"); err != nil {
		log.Println(color.RedString("ERROR"), "Failed to Fiber Listen")
		log.Fatalln(err.Error())
	}
}
