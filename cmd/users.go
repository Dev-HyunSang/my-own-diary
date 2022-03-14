package cmd

import (
	"log"
	"time"

	"github.com/dev-hyunsang/my-own-diary/config"
	"github.com/dev-hyunsang/my-own-diary/database"
	"github.com/dev-hyunsang/my-own-diary/model"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Logined
func HomeHandler(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv("SIGNING_KEY")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "로그인이 되지 않았어요. 다시 시도해 주세요.",
			"time":    time.Now(),
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user model.Users

	db, err := database.ConnectionDB()
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed to Connection DataBase")
		log.Println(err)
	}

	db.Where("uuid = ?", claims.Issuer).First(&user)

	return c.Status(200).JSON(fiber.Map{
		"stauts":  200,
		"message": "성공적으로 로그인 하였어요!",
		"data":    &user,
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
