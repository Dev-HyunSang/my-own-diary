package cmd

import (
	"log"
	"time"

	"github.com/dev-hyunsang/my-own-diary/database"
	"github.com/dev-hyunsang/my-own-diary/model"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/twinj/uuid"
)

func DiaryNewHandler(c *fiber.Ctx) error {
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
		user model.Users
	)

	req := new(model.Diary)

	db, err := database.ConnectionDB()
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed to Connection DataBase")
		log.Println(err)
	}

	db.Where("uuid =?", claims.Issuer).First(&user)

	diaryToken := uuid.NewV4()

	// 사용자가 입력한 비밀 일기장의 비밀번호를 암호화 함.
	diaryPassword, err := GeneratePassword(req.Password)
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed to Generate to Password")
		log.Println(err)
	}

	diary := model.Diary{
		DiaryUUID: diaryToken.String(),
		UserUUID:  user.UUID,
		Group:     req.Group,
		Content:   req.Content,
		Password:  diaryPassword,
		CreatedAt: time.Now(),
		RevisedAt: time.Now(),
	}

	db.Create(diary)

	// 추후 개선 필요
	// 현재는 ""의 경우에만 됨, 추후 프론트엔드 단과 백엔드단에서의 교차 검증 필요
	if req.Content == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "입력되지 않았어요. 다시 시도해 주세요.",
			"time":    time.Now(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"staus":   200,
		"message": "성공적으로 기록을 남겼어요.",
		"time":    time.Now().Format("2006-01-02 15:04:05"),
	})
}
