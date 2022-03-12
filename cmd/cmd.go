package cmd

import (
	"log"
	"time"

	"github.com/dev-hyunsang/my-own-diary/database"
	"github.com/dev-hyunsang/my-own-diary/model"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func IndexHandler(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello!",
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(hashVal, userPw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(userPw))
	if err != nil {
		return false
	} else {
		return true
	}
}

func RegisterHandler(c *fiber.Ctx) error {
	var user model.Register

	req := new(model.Register)
	if err := c.BodyParser(req); err != nil {
		log.Println(color.RedString("ERROR"), "Failed to Fiber BodyParser")
		log.Fatalln(err)
	}

	db, err := database.ConnectionDB()
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed Connection DataBase")
		log.Fatalln(err)
	}

	result := db.Find(&user, "email=?", user.Email)

	if result.RowsAffected != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"status":  "400",
			"message": "동일한 이메일 주소가 있어요, 다시 확인 해 주세요!",
			"time":    time.Now().String(),
		})
	}

	hashPw, err := HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{
			"status":  "500",
			"message": "비밀번호가 다릅니다.",
			"error":   err.Error(),
		})
	}
	user.Password = hashPw
	if err := db.Create(&user); err.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{
			"status":  "500",
			"message": "회원가입에 실패 하였어요. 다시 시도 해 주세요.",
			"time":    time.Now().String(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]string{
		"status":  "200",
		"message": "성공적으로 회원가입을 하였어요.",
		"time":    time.Now().String(),
	})
}
