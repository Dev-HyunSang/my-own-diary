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
		"uuid":  user.UUID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   exp,
	})

	// Create token
	token, err := claims.SignedString(config.GetEnv("SIGNING_KEY"))

	return token, err
}
