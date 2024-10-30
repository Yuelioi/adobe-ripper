package router

import (
	"adobe-ripper/internal/controllers"

	"github.com/gin-gonic/gin"
)

func Fire(r *gin.Engine) {
	system := r.Group("system")
	systemController := &controllers.SystemController{}
	system.GET("/ping", systemController.Pong)
	system.GET("/explore", systemController.Explore)
	system.GET("/trash", systemController.Trash)
}
