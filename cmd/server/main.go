package main

import (
	"fmt"
	"prrview/internal/config"
	"prrview/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", server.Health)

	if config.Env.AppEnv == config.AppEnvDevelopment {
		router.Run(fmt.Sprintf("localhost:%s", config.Env.Port))
	} else {
		router.Run(fmt.Sprintf(":%s", config.Env.Port))
	}

}
