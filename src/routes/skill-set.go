package routes

import (
	"github.com/KadirbekSharau/apprentice-backend/src/handlers"
	"github.com/KadirbekSharau/apprentice-backend/src/middleware"
	skillSetService "github.com/KadirbekSharau/apprentice-backend/src/services/skill-set"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Skill Set routes */
func InitSkillSetRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		repository = skillSetService.NewRepository(db)
		service    = skillSetService.NewService(repository)
		handler    = handlers.NewSkillSetHandler(service)
	)

	groupRoute := route.Group("/api/v1/skillset")
	groupRoute.POST("/", middleware.Auth([]int{0}), handler.CreateSkillSet)
	groupRoute.GET("/", handler.GetAllSkillSets)
	groupRoute.GET("/:id", handler.GetSkillSetByID)
}