package main

import (
	"log"
	"strings"

	"github.com/dev-hyunsang/my-own-diary/cmd"
	"github.com/dev-hyunsang/my-own-diary/database"
	"github.com/dev-hyunsang/my-own-diary/model"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db, err := database.ConnectionDB()
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed DataBase Connection")
		log.Fatalln(err.Error())
	}

	err = db.AutoMigrate(
		&model.Users{}, &model.Diary{})
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed to DataBase AutoMigrate")
		log.Fatalln(err)
	}
	log.Println(color.GreenString("SUCCESS"), "Successfully Create DataBase")

	app := fiber.New(fiber.Config{
		AppName: "my own diary",
	})
	app.Static("/", "./public")
	app.Use(cors.New(cors.Config{
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodDelete,
		}, ","),
	}))

	api := app.Group("/api")
	api.Get("/", cmd.IndexHandler)
	api.Post("/register", cmd.RegisterHandler)
	api.Post("/login", cmd.LoginHandler)

	user := app.Group("/api/user")
	user.Get("/home", cmd.HomeHandler)

	if err = app.Listen(":3000"); err != nil {
		log.Println(color.RedString("ERROR"), "Failed to Fiber Listen")
		log.Fatalln(err.Error())
	}
}
