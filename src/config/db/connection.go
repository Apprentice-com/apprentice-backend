package db

import (
	"log"

	"github.com/KadirbekSharau/apprentice-backend/src/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection() *gorm.DB {
	dbURL := "postgres://sharauq:sharauq@database:5432/apprentice"
	//dbURL := "host=localhost user=kadirbeksharau password=kadr2001 dbname=kadirbeksharau port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(
		&models.Users{},
		&models.Companies{},
		&models.Locations{},
		&models.CompanyImages{},
		&models.EducationDetails{},
		&models.EmployerProfiles{},
		&models.JobPosts{},
		&models.SeekerProfiles{},
		&models.SkillSets{},
		&models.UserLogs{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}
	AccountsDataMigrator(db)

	return db
}

func CloseDB(db *gorm.DB) {
	database, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	database.Close()
}
