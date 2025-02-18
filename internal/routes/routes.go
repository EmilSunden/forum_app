package routes

import (
	"app/internal/handlers"
	"app/internal/logger"
	"os"

	"app/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// Routes is the function that contains all the routes for the application
func Routes(db *gorm.DB) *gin.Engine {
	zeroLogLogger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	router := gin.Default()
	router.Use(logger.Logger(zeroLogLogger, db)) // create logger middleware that logs all requests
	authGroup := router.Group("/api/v1/auth")
	{
		authGroup.POST("/login", handlers.Login(db)) // au
		authGroup.POST("/signup", handlers.Signup(db))
		authGroup.POST("/logout", handlers.Logout)
	}

	protectedGroup := router.Group("/api/v1/protected")
	protectedGroup.Use(middleware.AuthMiddleware(db))
	{
		protectedGroup.POST("/friend-request", handlers.CreateFriendRequest(db))
		protectedGroup.POST("/friend-request/handle", handlers.HandleFriendRequest(db))
	}

	return router
}
