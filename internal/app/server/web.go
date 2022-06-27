package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nanoprobe/internal/app/configs"
)

func StartWebServer(cfg *configs.Config) error {
	router := gin.Default()
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "OK",
		})
	})
	return router.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
}
