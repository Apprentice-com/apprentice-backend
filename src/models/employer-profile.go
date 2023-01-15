package models

import (

	//"github.com/KadirbekSharau/bookswap-backend/util"
	"time"

	"gorm.io/gorm"
)

// CompanyProfiles is database entity for each transaction
type EmployerProfile struct {
	UserID      uint
	FirstName   string `gorm:"type:varchar(30)"`
	SecondName  string `gorm:"type:varchar(30)"`
	Description string `gorm:"type:varchar(1000)"`
}

// Company is database entity for categories
type Company struct {
	gorm.Model
	UserID             uint
	BusinessStreamID   uint
	CompanyName        string          `gorm:"type:varchar(100)"`
	CompanyDescription string          `gorm:"type:varchar(1000)"`
	EstablishmentDate  time.Time       `gorm:"type:date"`
	CompanyWebsiteUrl  string          `gorm:"type:varchar(500)"`
	CompanyJobs        []JobPost     `gorm:"foreignKey:CompanyID"`
	Images             []CompanyImage `gorm:"foreignKey:CompanyID"`
}

// CompanyImage is database entity for categories
type CompanyImage struct {
	gorm.Model
	CompanyID    uint   `gorm:"not null"`
	CompanyImage string `gorm:"type:varchar(500)"`
}

// Business Stream is a type of company
type BusinessStream struct {
	gorm.Model
	BusinessStreamName      string      `gorm:"not null"`
	BusinessStreamCompanies []Company `gorm:"foreignKey:BusinessStreamID"`
}
