package models

import "gorm.io/gorm"

// SeekerProfiles is database entity for each transaction
type SkillSets struct {
	gorm.Model
	Name             string            `gorm:"type:varchar(130)"`
	//SeekerSkillSets  []UserSkillSet  /* `gorm:"foreignKey:SkillSetID"` */ `gorm:"many2many:user_skill_sets;"`
	//JobPostSkillSets []JobPostSkillSet /* `gorm:"foreignKey:SkillSetID"` */ `gorm:"many2many:job_post_skill_sets;"`
}

// // SeekerSkillSet is database entity for each transaction
// type UserSkillSet struct {
// 	UserID     uint `gorm:"primaryKey"`
// 	SkillSetID uint `gorm:"primaryKey"`
// }

// // JobPostSkillSet is database entity for each transaction
// type JobPostSkillSet struct {
// 	JobPostID  uint `gorm:"primaryKey"`
// 	SkillSetID uint `gorm:"primaryKey"`
// }
