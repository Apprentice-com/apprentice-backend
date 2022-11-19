package models

import "gorm.io/gorm"

// SeekerProfiles is database entity for each transaction
type Locations struct {
	gorm.Model
	StreetAddress string     `gorm:"type:varchar(100)"`
	City          string     `gorm:"type:varchar(50)"`
	State         string     `gorm:"type:varchar(50)"`
	Country       string     `gorm:"type:varchar(50)"`
	ZipCode       string     `gorm:"type:varchar(50)"`
	JobPosts      []JobPosts `gorm:"foreignKey:LocationID"`
}
