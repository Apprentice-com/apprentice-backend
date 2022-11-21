package models

import "gorm.io/gorm"

// SeekerProfiles is database entity for each transaction
type JobPosts struct {
	gorm.Model
	UserID           uint
	CompanyID        uint
	LocationID       uint
	Name             string      `gorm:"type:varchar(1000)"`
	Description      string      `gorm:"type:varchar(130)"`
	IsActive         bool        `gorm:"default:true"`
	JobPostSkillSets []SkillSets `gorm:"many2many:job_post_skill_sets;"`
}
