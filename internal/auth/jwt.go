package auth

import (
	"app/internal/config"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID int64, username string) (string, error) {
	// Load secret key
	jwtConfig := config.LoadJWTConfigFromEnv()
	// Get the secret key
	secret := jwtConfig.GetJWTSecret()

	// Set token expiration appropriately
	expirationTime := time.Now().Add(24 * time.Hour).Unix()

	// Create the Claims
	claims := jwt.MapClaims{
		"iss":      "stream-mix",
		"sub":      strconv.FormatInt(userID, 10),
		"username": username,
		"exp":      expirationTime,
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(secret))
	return signedToken, nil
}

func ValidateJWTAndExtractClaims(tokenString string) (jwt.MapClaims, error) {
	// Load secret key
	jwtConfig := config.LoadJWTConfigFromEnv()
	// Get the secret key
	secret := jwtConfig.GetJWTSecret()
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}
	// Extract and validate the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
