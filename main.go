package main

import (
	"github.com/KadirbekSharau/apprentice-backend/src/config/db"
	"github.com/KadirbekSharau/apprentice-backend/src/routes"
	"github.com/KadirbekSharau/apprentice-backend/src/util"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db.NewDatabaseConnection()
	server := gin.Default()

	if util.GodotEnv("GO_ENV") == "debug" {
		gin.SetMode(gin.DebugMode)
	} 
	if util.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} 
	if util.GodotEnv("GO_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	server.Use(
		cors.New(cors.Config{
			AllowOrigins:  []string{"*"},
			AllowMethods:  []string{"*"},
			AllowHeaders:  []string{"*"},
			AllowWildcard: true,
		}),
	)

	server.Use(helmet.Default())
	server.Use(gzip.Gzip(gzip.BestCompression))

	routes.InitAuthRoutes(db, server)
	routes.InitProfileRoutes(db, server)
	routes.InitSkillSetRoutes(db, server)
	routes.InitJobPostRoutes(db, server)

	server.Run(":" + util.GodotEnv("GO_PORT"))
}
