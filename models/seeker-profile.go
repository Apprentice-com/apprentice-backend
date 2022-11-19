package models

import "time"

// SeekerProfiles is database entity for each transaction
type SeekerProfiles struct {
	UserID      uint 
	FirstName   string `gorm:"type:varchar(30)"`
	SecondName  string `gorm:"type:varchar(30)"`
	Description string `gorm:"type:varchar(1000)"`
	Resume      string `gorm:"type:varchar"`
}

// EducationDetails is database entity for each transaction
type EducationDetails struct {
	UserID          uint
	Degree          string    `gorm:"primaryKey; type:varchar(50)"`
	Major           string    `gorm:"primaryKey; type:varchar(100)"`
	InstitutionName string    `gorm:"primaryKey; type:varchar"`
	StartDate       time.Time `gorm:"type:date"`
	EndDate         time.Time `gorm:"type:date"`
}
