// JWT handling code
package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yuvakkrishnan/user-activity-logger/pkg/models"
)

// TODO: Implement JWT functions

func GenerateJwt(userID string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(models.Jwtkey)
}
func ValidateJWT(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return models.Jwtkey, nil
	})
	return err == nil && token.Valid
}
