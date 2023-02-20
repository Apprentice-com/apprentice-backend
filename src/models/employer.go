package models

import (
	"gorm.io/gorm"
	"time"
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
	CompanyName        string    `gorm:"type:varchar(100)"`
	CompanyDescription string    `gorm:"type:varchar(1000)"`
	EstablishmentDate  time.Time `gorm:"type:date"`
	CompanyWebsiteUrl  string    `gorm:"type:varchar(500)"`
	CompanyJobs        []JobPost `gorm:"foreignKey:CompanyID"`
	CompanyImage       string    `gorm:"type:varchar(500)"`
}
