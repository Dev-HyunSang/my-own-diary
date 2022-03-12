package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/dev-hyunsang/my-own-diary/database"
	"github.com/dev-hyunsang/my-own-diary/model"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hash), err
}

func RegisterHandler(c *fiber.Ctx) error {
	req := new(model.Register)
	if err := c.BodyParser(req); err != nil {
		log.Println(color.RedString("ERROR"), "Failed to Fiber BodyParser")
		log.Fatalln(err)
	}

	db, err := database.ConnectionDB()
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed to Connection DataBase")
		log.Fatalln(err)
	}

	hashPw, err := GeneratePassword(req.Password)
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed GenerateForm Password")
		log.Println(err)
	}

	userUUID := uuid.NewV4()
	data := model.Users{
		UUID:      userUUID,
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashPw,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&data)
	log.Println(color.GreenString("SUCCESS"), fmt.Sprintf("UUID:%s | 생성 완료.", data.UUID))

	return c.Status(200).JSON(map[string]string{
		"stauts":  "200",
		"message": fmt.Sprintf("%s님 어서오세요! 나만의 일기에 오신 것을 환영해요!", req.Name),
		"time":    time.Now().String(),
	})
}

func LoginIndexHandler(c *fiber.Ctx) error {
	var data *model.Users
	req := new(model.Login)
	if err := c.BodyParser(req); err != nil {
		log.Println(color.RedString("ERROR"), "Failed to BodyParser")
		log.Fatalln(err)
	}

	db, err := database.ConnectionDB()
	if err != nil {
		log.Println(color.RedString("ERRPR"), "Failed to Connection DataBase")
		log.Fatalln(err)
	}
	result := db.Where("email = ?", req.Email).Find(&data)

	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "입력해 주신 정보로 회원 정보를 찾을 수 없네요. 다시 시도 해 주세요.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"status":  "400",
			"message": "입력해 주신 정보로 회원 정보를 찾을 수 없네요. 다시 시도 해 주세요.",
			"time":    time.Now().String(),
		})
	}

	token, exp, err := CreateJWT(data)
	if err != nil {
		log.Println(color.RedString("ERROR", "Failed to Create Json Web Token"))
		log.Println(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "로그인을 성공적으로 했어요!",
		"token":   token,
		"exp":     exp,
		"user":    data,
		"time":    time.Now(),
	})
}
