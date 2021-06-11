package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rjunior/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("about to start the application...")

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
