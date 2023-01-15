package models

import "gorm.io/gorm"

// JobPost is database entity for each job post
type JobPost struct {
	gorm.Model
	UserID           uint
	CompanyID        uint
	LocationID       uint
	JobPostTypeID    uint
	IsRemote         bool
	IsPaid           bool
	Salary           int
	Currency         string
	Name             string     `gorm:"type:varchar(1000)"`
	Description      string     `gorm:"type:varchar(130)"`
	Link             string     `gorm:"type:varchar(1000)"`
	IsActive         bool       `gorm:"default:true"`
	JobPostSkillSets []SkillSet `gorm:"many2many:job_post_skill_sets;"`
}

// JobPostType is database entity for each job post type
type JobPostType struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(1000)"`
	JobPosts []JobPost `gorm:"foreignKey:JobPostTypeID"`
}
