package routes

import (
	"app/internal/auth"
	"app/internal/friends"

	"app/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Routes is the function that contains all the routes for the application
func Routes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	authGroup := router.Group("/api/v1/auth")
	{
		authGroup.POST("/login", auth.Login(db))
		authGroup.POST("/signup", auth.Signup(db))
		authGroup.POST("/logout", auth.Logout)
	}

	protectedGroup := router.Group("/api/v1/protected")
	protectedGroup.Use(middleware.AuthMiddleware(db))
	{
		protectedGroup.POST("/friend-request", friends.CreateFriendRequest(db))
		protectedGroup.POST("/friend-request/handle", friends.HandleFriendRequest(db))
	}

	return router
}
