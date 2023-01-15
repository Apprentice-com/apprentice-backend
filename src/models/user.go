package models

import (
	"time"

	"github.com/KadirbekSharau/apprentice-backend/src/util"
	"gorm.io/gorm"
)

// Users is database entity for user
type Users struct {
	gorm.Model
	UserType        int
	Email           string            `gorm:"type:varchar(50);unique;not null"`
	Password        string            `gorm:"type:varchar(255)"`
	IsActive        bool              `gorm:"type:bool"`
	ContactNumber   string            `gorm:"type:varchar(20)"`
	UserImage       string            `gorm:"type:varchar(255)"`
	Companies       []Company         `gorm:"foreignKey:UserID"`
	Degrees         []EducationDetail `gorm:"foreignKey:UserID"`
	SeekerSkillSets []SkillSet        `gorm:"many2many:user_skill_sets;"`
	Logs            []UserLog         `gorm:"foreignKey:UserID"`
	Jobposts        []JobPost         `gorm:"foreignKey:UserID"`
	SeekerProfile   []SeekerProfile   `gorm:"foreignKey:UserID"`
	EmployerProfile []EmployerProfile `gorm:"foreignKey:UserID"`
}

type UserLog struct {
	UserID           uint      `gorm:"not null"`
	LastLoginDate    time.Time `gorm:"type:date; not null"`
	LastJobApplyDate time.Time `gorm:"type:date; not null"`
}

func (entity *Users) BeforeCreate(db *gorm.DB) error {
	entity.Password = util.HashPassword(entity.Password)
	return nil
}
