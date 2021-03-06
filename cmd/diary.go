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

func NewDiaryHandler(c *fiber.Ctx) error {
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
	if err := c.BodyParser(req); err != nil {
		log.Println(color.RedString("ERROR"), "Failed to BodyParser")
		log.Println(err)
	}

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
		Title:     req.Title,
		Group:     req.Group,
		Content:   req.Content,
		Password:  diaryPassword,
		CreatedAt: time.Now(),
		RevisedAt: time.Now(),
	}

	result := db.Create(&diary)
	log.Println(color.GreenString("SUCCESS"), "Created Diary Content", result.RowsAffected)

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

func AllDiaryListHandler(c *fiber.Ctx) error {
	var (
		diaryData []model.Diary
	)

	cookie := c.Cookies("jwt")
	claims, err := VerificationJWT(cookie)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to Verification JWT",
			"time":    time.Now(),
		})
	}

	db, err := database.ConnectionDB()
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed to Connection DataBase")
		log.Println(err)
	}

	reslut := db.Where("user_uuid = ?", claims.Issuer).Find(&diaryData)
	if reslut.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  200,
			"message": "아직 등록되어 있는 일기가 없어요! 저랑 같이 일기 쓰시러 가실래요?",
			"time":    time.Now(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "성공적으로 모든 일기장을 불러왔어요!",
		"datas":   diaryData,
		"time":    time.Now(),
	})
}
