package main

import (
	"log"

	"github.com/dev-hyunsang/my-own-diary/database"
	"github.com/dev-hyunsang/my-own-diary/model"
	"github.com/dev-hyunsang/my-own-diary/router"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "my own diary",
	})

	router.Middleware(app)

	db, err := database.ConnectionDB()
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed DataBase Connection")
		log.Fatalln(err.Error())
	}

	err = db.AutoMigrate(model.Login{}, model.Register{})
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed to DataBase AutoMigrate")
		log.Fatalln(err)
	}

	if err = app.Listen(":3000"); err != nil {
		log.Println(color.RedString("ERROR"), "Failed to Fiber Listen")
		log.Fatalln(err.Error())
	}
}
