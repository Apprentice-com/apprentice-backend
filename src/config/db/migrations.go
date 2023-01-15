package db

import (
	"log"
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"github.com/KadirbekSharau/apprentice-backend/src/models"
	authService "github.com/KadirbekSharau/apprentice-backend/src/services/auth"
	"gorm.io/gorm"
)

func accountsDataMigrator(db *gorm.DB) *models.Users {
	registerRepository := authService.NewRepository(db)
	registerService := authService.NewService(registerRepository)
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

func businessMigrator(db *gorm.DB) (*[]models.BusinessStream, int, string) {
	for _, bus := range businessStreams {
		if db.Debug().Select("*").Where("business_stream_name = ?", bus.BusinessStreamName).Find(&models.BusinessStream{}).RowsAffected > 0 {
			return nil, http.StatusConflict, "Already exists"
		}
		if db.Debug().Create(&bus).Error != nil {
			return nil, http.StatusForbidden, "Create Failed"
		}
	}
	db.Commit()
	return &businessStreams, http.StatusOK, "nil"
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

func skillSetMigrator(db *gorm.DB) (*[]models.SkillSet, int, string) {
	for _, skillset := range skillSets {
		if db.Debug().Select("*").Where("name = ?", skillset.Name).Find(&models.SkillSet{}).RowsAffected > 0 {
			return nil, http.StatusConflict, "Already exists"
		}
		if db.Debug().Create(&skillset).Error != nil {
			return nil, http.StatusForbidden, "Create Failed"
		}
	}
	db.Commit()
	return &skillSets, http.StatusOK, "nil"
}

func jobPostsMigrator(db *gorm.DB) (*[]models.JobPost, int, string) {
	for _, jobPost := range jobPosts {
		if db.Debug().Create(&jobPost).Error != nil {
			return nil, http.StatusForbidden, "Create Failed"
		}
	}
	db.Commit()
	return &jobPosts, http.StatusOK, "nil"
}
