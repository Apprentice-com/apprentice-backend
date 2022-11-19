package models

import (

	//"github.com/KadirbekSharau/bookswap-backend/util"
	"time"

	"gorm.io/gorm"
)

// CompanyProfiles is database entity for each transaction
type EmployerProfiles struct {
	UserID      uint
	FirstName   string `gorm:"type:varchar(30)"`
	SecondName  string `gorm:"type:varchar(30)"`
	Description string `gorm:"type:varchar(1000)"`
}

// Companies is database entity for categories
type Companies struct {
	gorm.Model
	UserID             uint
	CompanyName        string          `gorm:"type:varchar(100)"`
	CompanyDescription string          `gorm:"type:varchar(1000)"`
	EstablishmentDate  time.Time       `gorm:"type:date"`
	CompanyWebsiteUrl  string          `gorm:"type:varchar(500)"`
	CompanyJobs        []JobPosts      `gorm:"foreignKey:CompanyID"`
	Images             []CompanyImages `gorm:"foreignKey:CompanyID"`
}

// CompanyImages is database entity for categories
type CompanyImages struct {
	gorm.Model
	CompanyID    uint   `gorm:"not null"`
	CompanyImage string `gorm:"type:varchar(500)"`
}
