package auth

import (
	"TRANSFERSYSTEM/app/config"
	"TRANSFERSYSTEM/model"

	"github.com/golang-jwt/jwt"
)

func NewToken(user model.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     user.ID,
		"name":    user.Name,
	})

	signedToken, _ := token.SignedString([]byte(config.Get("SECRET")))

	return signedToken
}
