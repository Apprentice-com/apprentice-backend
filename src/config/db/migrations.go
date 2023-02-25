package db

import (
	"log"
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"github.com/KadirbekSharau/apprentice-backend/src/models"
	"github.com/KadirbekSharau/apprentice-backend/src/services"
	"gorm.io/gorm"
)

func accountsDataMigrator(db *gorm.DB) *models.Users {
	registerService := services.NewUserService(models.NewUserRepository(db))
	admin := dto.InputUserSeekerRegister{
		FirstName:  "Admin1",
		SecondName: "Admin1ov",
		Email:      "admin1@gmail.com",
		Password:   "admin532",
	}
	data, status, err := registerService.AdminRegister(&admin)
	if status != 201 {
		log.Println(err)
	}
	return data
}

func jobPostTypeMigrator(db *gorm.DB) (*[]models.JobPostType, int, string) {
	for _, jobType := range jobTypes {
		if db.Debug().Select("*").Where("name = ?", jobType.Name).Find(&models.JobPostType{}).RowsAffected > 0 {
			return nil, http.StatusConflict, "Already exists"
		}
		if db.Debug().Create(&jobType).Error != nil {
			return nil, http.StatusForbidden, "Create Failed"
		}
	}
	db.Commit()
	return &jobTypes, http.StatusOK, "nil"
}

func locationMigrator(db *gorm.DB) (*[]models.Location, int, string) {
	for _, location := range locations {
		if db.Debug().Select("*").Where("city = ?", location.City).Find(&models.Location{}).RowsAffected > 0 {
			return nil, http.StatusConflict, "Already exists"
		}
		if db.Debug().Create(&location).Error != nil {
			return nil, http.StatusForbidden, "Create Failed"
		}
	}
	db.Commit()
	return &locations, http.StatusOK, "nil"
}

func companyMigrator(db *gorm.DB) (*[]models.Company, int, string) {
	for _, company := range companies {
		if db.Debug().Select("*").Where("company_name = ?", company.CompanyName).Find(&models.Company{}).RowsAffected > 0 {
			return nil, http.StatusConflict, "Already exists"
		}
		if db.Debug().Create(&company).Error != nil {
			return nil, http.StatusForbidden, "Create Failed"
		}
	}
	db.Commit()
	return &companies, http.StatusOK, "nil"
}

func jobPostsMigrator(db *gorm.DB) (*[]models.JobPost, int, string) {
	if db.Model(&models.JobPost{}).Debug().Select("*").Find(&models.JobPost{}).RowsAffected > 0 {
		return nil, http.StatusConflict, "Already exists"
	}
	for _, jobPost := range jobPosts {
		if db.Debug().Create(&jobPost).Error != nil {
			return nil, http.StatusForbidden, "Create Failed"
		}
	}
	db.Commit()
	return &jobPosts, http.StatusOK, "nil"
}
