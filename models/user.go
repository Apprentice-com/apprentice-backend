package models

import (
	"time"

	"github.com/KadirbekSharau/apprentice-backend/util"
	"gorm.io/gorm"
)

// Users is database entity for user
type Users struct {
	gorm.Model
	UserType        string             `gorm:"not null"`
	Email           string             `gorm:"type:varchar(50);unique;not null"`
	Password        string             `gorm:"type:varchar(255)"`
	IsActive        bool               `gorm:"type:bool"`
	ContactNumber   string             `gorm:"type:varchar(20);unique;not null"`
	UserImage       string             `gorm:"type:varchar(255)"`
	Companies       []Companies        `gorm:"foreignKey:UserID"`
	Degrees         []EducationDetails `gorm:"foreignKey:UserID"`
	SeekerSkillSets []SkillSets        `gorm:"many2many:user_skill_sets;"`
	Logs            []UserLogs         `gorm:"foreignKey:UserID"`
	Jobposts        []JobPosts         `gorm:"foreignKey:UserID"`
	SeekerProfile   []SeekerProfiles   `gorm:"foreignKey:UserID"`
	EmployerProfile []EmployerProfiles `gorm:"foreignKey:UserID"`
}

type UserLogs struct {
	UserID           uint      `gorm:"not null"`
	LastLoginDate    time.Time `gorm:"type:date; not null"`
	LastJobApplyDate time.Time `gorm:"type:date; not null"`
}

func (entity *Users) BeforeCreate(db *gorm.DB) error {
	entity.Password = util.HashPassword(entity.Password)
	return nil
}
