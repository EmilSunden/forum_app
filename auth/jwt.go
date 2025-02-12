package auth

import (
	"app/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user string) (string, error) {
	// Load secret key
	jwtConfig := config.LoadJWTConfigFromEnv()
	// Get the secret key
	secret := jwtConfig.GetJWTSecret()
	// Create the Claims
	claims := jwt.MapClaims{
		"iss": "stream-mix",
		"sub": user,
		"exp": 15000,
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(secret))
	return signedToken, nil
}

func ValidateJWT(tokenString string) bool {
	// Load secret key
	jwtConfig := config.LoadJWTConfigFromEnv()
	// Get the secret key
	secret := jwtConfig.GetJWTSecret()
	// Parse the token
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	// Check if the token is valid
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}
	return false

}
