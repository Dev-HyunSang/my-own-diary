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
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.UUID,         // 로그인한 사용자의 UUID를 기록함.
		IssuedAt:  time.Now().Unix(), // time.Now을 사용할 경우 time.time으로 반환됨, unix를 사용하는 경우 int64로 반환함.
		ExpiresAt: exp,
	})

	// Create token
	token, err := claims.SignedString([]byte(config.GetEnv("SIGNING_KEY")))

	return token, err
}

func VerificationJWT(cookie string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv("SIGNING_KEY")), nil
	})

	clamis := token.Claims.(*jwt.StandardClaims)

	return clamis, err
}
