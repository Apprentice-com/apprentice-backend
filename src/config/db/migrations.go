package db

import (
	"log"

	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"github.com/KadirbekSharau/apprentice-backend/src/models"
	authService "github.com/KadirbekSharau/apprentice-backend/src/services/auth"
	"gorm.io/gorm"
)

func AccountsDataMigrator(db *gorm.DB) (*models.Users) {
	registerRepository := authService.NewRepository(db)
	registerService := authService.NewService(registerRepository)
	admin := dto.InputUserSeekerRegister{
		FirstName: "Admin1",
		SecondName: "Admin1ov",
		Email: "admin1@gmail.com",
		Password: "admin532",
	}
	newAdmin, errAdmin := registerService.AdminRegister(&admin)
	if errAdmin == "REGISTER_CONFLICT_409" || errAdmin == "REGISTER_FAILED_403" {
		log.Println(errAdmin)
	}

	return newAdmin;
}