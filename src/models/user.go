package models

import (
	"fmt"
	"net/http"
	"time"

	"github.com/KadirbekSharau/apprentice-backend/src/dto"
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
	Jobposts        []JobPost         `gorm:"foreignKey:UserID"`
	SeekerProfile   []SeekerProfile   `gorm:"foreignKey:UserID"`
	EmployerProfile []EmployerProfile `gorm:"foreignKey:UserID"`
}

func (entity *Users) BeforeCreate(db *gorm.DB) error {
	entity.Password = util.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Users) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

/* User Login Repository Service */
func (r *UserRepository) UserLogin(input *dto.InputLogin) (*Users, int, string) {
	var users Users
	db := r.db.Model(&users)
	users.Email = input.Email
	users.Password = input.Password

	if db.Debug().Select("*").Where("email = ?", input.Email).Find(&users).RowsAffected < 1 {
		return nil, http.StatusNotFound, "User account is not registered"
	}
	if !users.IsActive {
		return nil, http.StatusForbidden, "User account is not active"
	}
	if util.ComparePassword(users.Password, input.Password) != nil {
		fmt.Println(users.Password)
		fmt.Println(input.Password)
		fmt.Println(util.ComparePassword(users.Password, input.Password))
		return nil, http.StatusForbidden, "Password is wrong"
	}
	return &users, http.StatusOK, "Logged in successfully"
}

/* User Registration Repository */
func (r *UserRepository) UserRegister(isActive bool, userType int, input *dto.InputUserSeekerRegister) (*Users, int, string) {
	var users Users
	db := r.db.Model(&users)
	if db.Debug().Select("*").Where("email = ?", input.Email).Find(&users).RowsAffected > 0 {
		return nil, http.StatusConflict, "Email already exists"
	}
	users.Email = input.Email
	users.Password = input.Password
	users.IsActive = isActive
	users.UserType = userType
	if db.Debug().Create(&users).Error != nil {
		return nil, http.StatusForbidden, "Registering new account failed"
	}
	db.Commit()
	if userType == 1 {
		r.AddNewSeekerProfile(users.ID, input)
	}
	return &users, http.StatusCreated, "Registered successfully"
}

/* Adding Seeker Profile Repository */
func (r *UserRepository) AddNewSeekerProfile(userId uint, input *dto.InputUserSeekerRegister) (*SeekerProfile, string) {
	var seekerProfile SeekerProfile
	db := r.db.Model(&seekerProfile)

	if db.Debug().Select("*").Where("user_id = ?", userId).Find(&seekerProfile).RowsAffected > 0 {
		return nil, "REGISTER_CONFLICT_409"
	}

	seekerProfile.UserID = userId
	seekerProfile.FirstName = input.FirstName
	seekerProfile.SecondName = input.SecondName

	if db.Debug().Create(&seekerProfile).Error != nil {
		return nil, "REGISTER_FAILED_403"
	}
	db.Commit()
	return &seekerProfile, "nil"
}
