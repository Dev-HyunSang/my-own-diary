package cmd

import (
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/dev-hyunsang/my-own-diary/config"
	"github.com/dev-hyunsang/my-own-diary/model"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func SendEmail(email string) {
	auth := smtp.PlainAuth("", config.GetEnv("MAIL_ADDRESS"), config.GetEnv("SMTP_PASSWORD"), config.GetEnv("SMTP_SERVER_ADDRESS"))

	from := config.GetEnv("MAIL_ADDRESS")
	to := []string{email}

	headerSubject := "Subject: 테스트\r\n"
	headerBlank := "\r\n"
	body := "HeckingEmail 테스트입니당 ~\r\n"
	msg := []byte(headerSubject + headerBlank + body)

	err := smtp.SendMail(fmt.Sprintf("%s:%s", config.GetEnv("SMTP_SERVER_ADDRESS"), config.GetEnv("SMTP_SERVER_PORT")), auth, from, to, msg)
	if err != nil {
		log.Println(color.RedString("ERROR"), "Failed to Send Email")
		log.Println(err)
	}
	log.Println(color.GreenString("SUCCESS"), "Successfully to Send Email!")

}

func TestSendMailHandler(c *fiber.Ctx) error {
	req := new(model.Email)
	if err := c.BodyParser(req); err != nil {
		log.Println(color.RedString("ERROR"), "Failed to Fiber BodyParser")
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "오류가 발생했어요. 잠시후 다시 시도 해 주세요.",
			"time":    time.Now(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "성공적으로 메일을 보냈어요!",
		"time":    time.Now(),
	})
}
