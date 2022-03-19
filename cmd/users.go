package cmd

import (
	"log"
	"time"

	"github.com/dev-hyunsang/my-own-diary/database"
	"github.com/dev-hyunsang/my-own-diary/model"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

// Logined
func HomeHandler(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	claims, err := VerificationJWT(cookie)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to Verification JWT",
			"time":    time.Now(),
		})
	}

	var (
		diary []model.Diary
		user  model.Users
	)

	db, err := database.ConnectionDB()
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed to Connection DataBase")
		log.Println(err)
	}

	db.Where("uuid = ?", claims.Issuer).First(&user)

	db.Where("user_uuid = ?", claims.Issuer).Find(&diary)

	return c.Status(200).JSON(fiber.Map{
		"stauts":  200,
		"message": "성공적으로 로그인 하였어요!",
		"data":    &user,
		"diary":   &diary,
		"time":    time.Now(),
	})
}

// LogOut
func LogOutHandler(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
