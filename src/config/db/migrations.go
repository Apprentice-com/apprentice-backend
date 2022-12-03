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
	data, status, err := registerService.AdminRegister(&admin)
	if status != 201 {
		log.Println(err)
	}

	return data;
}