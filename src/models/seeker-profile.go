package models

import "time"

// SeekerProfile is database entity for each transaction
type SeekerProfile struct {
	UserID      uint 
	FirstName   string `gorm:"type:varchar(30)"`
	SecondName  string `gorm:"type:varchar(30)"`
	Description string `gorm:"type:varchar(1000)"`
	Resume      string `gorm:"type:varchar"`
}

// EducationDetail is database entity for each transaction
type EducationDetail struct {
	UserID          uint
	Degree          string    `gorm:"primaryKey; type:varchar(50)"`
	Major           string    `gorm:"primaryKey; type:varchar(100)"`
	InstitutionName string    `gorm:"primaryKey; type:varchar"`
	StartDate       time.Time `gorm:"type:date"`
	EndDate         time.Time `gorm:"type:date"`
}
