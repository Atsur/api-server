package api

import (
	"fmt"
	"log"

	"github.com/atsur/api-server/internal/api/router"
	"github.com/atsur/api-server/internal/pkg/config"
	"github.com/atsur/api-server/internal/pkg/db"
	"github.com/atsur/api-server/pkg/fbauth"
	"github.com/gin-gonic/gin"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "data/config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	client, err := fbauth.InitAuth()
	if err != nil {
		log.Fatalln("failed to init firebase auth", err)
	}
	web := router.Setup(client)
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
