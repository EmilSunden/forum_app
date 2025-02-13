package middleware

import (
	"app/auth"
	contextKeys "app/contextkeys"
	"app/internal/services"
	"context"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the token from the cookie
			cookie, err := r.Cookie("token")
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// assign the value of the cookie which is the token to the token variable
			tokenString := cookie.Value

			// Validate the token
			claims, err := auth.ValidateJWTAndExtractClaims(tokenString)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Extract user ID from claims (assuming "sub" holds the user ID as a string)
			sub, ok := claims["sub"].(string)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userID, err := strconv.ParseInt(sub, 10, 64)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Get the user from the database using the user ID
			user, err := services.GetUserByID(db, userID)
			if err != nil || user == nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Store the user in context using our defined key.
			ctx := context.WithValue(r.Context(), contextKeys.UserKey, user)

			// Token is valid, proceed to the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
