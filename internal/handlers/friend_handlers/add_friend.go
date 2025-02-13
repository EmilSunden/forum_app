package friendhandlers

import (
	contextKeys "app/contextkeys"
	"app/dtos"
	"app/internal/models"
	"app/internal/repositories"
	"app/internal/services"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func CreateFriendRequest(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// retrieve the current user from the requests's context.
		user, ok := r.Context().Value(contextKeys.UserKey).(*models.User)
		if !ok || user == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Parse the incoming JSON payload into DTO
		var input dtos.CreateFriendRequestInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		repo := repositories.NewFriendRequestRepository(db)
		friendService := services.FriendRequestService{Repo: repo}
		friendRequest := models.FriendRequest{

			ReceiverID: input.ReceiverID,
			Message:    input.Message,
		}

		// Use the service layer to create a friend request.

		if err := friendService.CreateFriendRequest(int64(user.ID), friendRequest); err != nil {
			http.Error(w, "Error creating friend request", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(map[string]string{"message": "Friend request created successfully"})
	}
}
