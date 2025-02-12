package auth

import (
	"app/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user string) (string, error) {
	// Get secret key
	secret := config.LoadJWTSecretFromEnv().GetJWTSecret()
	// Create the Claims
	claims := jwt.MapClaims{
		"iss": "streaming_forum_app",
		"sub": user,
		"exp": 15000,
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(secret))
	return signedToken, nil
}

func ValidateJWT(tokenString string) bool {
	secret := config.LoadJWTSecretFromEnv().GetJWTSecret()
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}
	return false

}
