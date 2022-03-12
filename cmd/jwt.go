package cmd

import (
	"time"

	"github.com/dev-hyunsang/my-own-diary/config"
	"github.com/dev-hyunsang/my-own-diary/model"
	"github.com/golang-jwt/jwt/v4"
)

func CreateJWT(user *model.Users) (string, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()

	// Create the Claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":      user.UUID.String(), // 로그인한 사용자의 UUID를 기록함.
		"IssuedAt":  time.Now().Unix(),  // time.Now을 사용할 경우 time.time으로 반환됨, unix를 사용하는 경우 int64로 반환함.
		"ExpiresAt": exp,
	})

	// Create token
	token, err := claims.SignedString(config.GetEnv("SIGNING_KEY"))

	return token, err
}
