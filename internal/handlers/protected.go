package handlers

import (
	"encoding/json"
	"net/http"
)

func ProtectedMux() http.Handler {
	// ProtectedMux is the function that contains the protected routes for the application
	// Protected routes logic here
	protectedMux := http.NewServeMux()
	protectedMux.Handle("/profile", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		response := map[string]string{"message": "This is the profile endpoint"}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(response)
	}))

	return http.StripPrefix("/api/v1/protected", protectedMux)
}
