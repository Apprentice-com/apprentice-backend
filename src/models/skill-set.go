package models

import "gorm.io/gorm"

// SeekerProfiles is database entity for each transaction
type SkillSets struct {
	gorm.Model
	Name             string            `gorm:"type:varchar(130); unique"`
}