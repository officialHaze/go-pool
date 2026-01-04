package server

import (
	"fmt"
	"log"
	"gopool/api/REST/server/routes"
	"gopool/settings"
	"gopool/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	if !util.InDevMode() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Cors setup
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Type"},
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
	}))

	// Map routes
	routes.MapRoutes(r)

	addr := fmt.Sprintf(":%d", settings.MySettings.SERVER_PORT)
	if err := r.Run(addr); err != nil {
		log.Fatalln(err)
	}
}
