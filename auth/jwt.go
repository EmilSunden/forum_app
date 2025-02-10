package auth

import (
	"app/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func generateJWT(user string) string {
	// Get secret key
	secret := config.LoadJWTSecretFromEnv().GetJWTSecret()
	// Create the Claims
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = user

	token := jwt.New(jwt.SigningMethodHS256)
	signedToken, _ := token.SignedString([]byte(secret))
	return signedToken
}

func validateJWT(tokenString string) bool {
	secret := config.LoadJWTSecretFromEnv().GetJWTSecret()
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}
	return false

}
