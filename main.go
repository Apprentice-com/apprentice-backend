package main

import (
	"github.com/KadirbekSharau/apprentice-backend/config/db"
	"github.com/KadirbekSharau/apprentice-backend/routes"
	"github.com/KadirbekSharau/apprentice-backend/util"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db.NewDatabaseConnection()
	server := gin.Default()

	if util.GodotEnv("GO_ENV") != "production" && util.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if util.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
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
	// routes.InitCategoryRoutes(db, server)

	server.Run(":8080")
}