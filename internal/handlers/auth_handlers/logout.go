package auth_handlers

import (
	"net/http"
	"time"
)

func Logout() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the token cookie with an expired date to effectively log the user out
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    "",
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
			Expires:  time.Unix(0, 0),
		})
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Logged out successfully"))
	})
}
