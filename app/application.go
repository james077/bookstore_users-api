package app

import (
	"github.com/gin-gonic/gin"
	"github.com/james077/bookstore_utils-go/logger"
)

var(
	router = gin.Default()
)
// StartApplication a
func StartApplication(){
	mapUrls()
	logger.Info("Inicio de aplicación...")

	router.Run(":8081")
}