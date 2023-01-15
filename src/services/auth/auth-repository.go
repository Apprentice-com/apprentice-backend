package authService

import (
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	model "github.com/KadirbekSharau/apprentice-backend/src/models"
	util "github.com/KadirbekSharau/apprentice-backend/src/util"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

/* User Login Repository Service */
func (r *repository) UserLogin(input *dto.InputLogin) (*model.Users, int, string) {
	var users model.Users
	db := r.db.Model(&users)
	users.Email = input.Email
	users.Password = input.Password

	if db.Debug().Select("*").Where("email = ?", input.Email).Find(&users).RowsAffected < 1 {
		return &users, http.StatusNotFound, "User account is not registered"
	}
	if !users.IsActive {
		return &users, http.StatusForbidden, "User account is not active"
	}
	if util.ComparePassword(users.Password, input.Password) != nil {
		return &users, http.StatusForbidden, "Password is wrong"
	}
	return &users, http.StatusOK, "Logged in successfully"
}

/* Active User Seeker Registration Repository */
func (r *repository) ActiveUserSeekerRegister(input *dto.InputUserSeekerRegister) (*model.Users, int, string) {
	var users model.Users
	db := r.db.Model(&users)
	if db.Debug().Select("*").Where("email = ?", input.Email).Find(&users).RowsAffected > 0 {
		return &users, http.StatusConflict, "Email already exists"
	}
	users.Email = input.Email
	users.Password = input.Password
	users.IsActive = true
	users.UserType = 1
	if db.Debug().Create(&users).Error != nil {
		return nil, http.StatusForbidden, "Registering new account failed"
	}
	db.Commit()
	r.AddNewSeekerProfile(users.ID, input)
	return &users, http.StatusCreated, "Registered successfully"
}

/* Active User Employer Registration Repository */
func (r *repository) ActiveUserEmployerRegister(input *dto.InputUserSeekerRegister) (*model.Users, int, string) {
	var users model.Users
	db := r.db.Model(&users)

	if db.Debug().Select("*").Where("email = ?", input.Email).Find(&users).RowsAffected > 0 {
		return &users, http.StatusConflict, "Email already exists"
	}

	// Need to create a employee profile
	users.Email = input.Email
	users.Password = input.Password
	users.IsActive = true
	users.UserType = 2

	if db.Debug().Create(&users).Error != nil {
		return nil, http.StatusForbidden, "Registering new account failed"
	}
	db.Commit()
	return &users, http.StatusCreated, "Registered successfully"
}

/* Admin Registration Repository */
func (r *repository) AdminRegister(input *dto.InputUserSeekerRegister) (*model.Users, int, string) {
	var users model.Users
	db := r.db.Model(&users)

	if db.Debug().Select("*").Where("email = ?", input.Email).Find(&users).RowsAffected > 0 {
		return &users, http.StatusConflict, "Email already exists"
	}

	// Need to create a employee profile
	users.Email = input.Email
	users.Password = input.Password
	users.IsActive = true
	users.UserType = 0

	if db.Debug().Create(&users).Error != nil {
		return nil, http.StatusForbidden, "Registering new account failed"
	}
	db.Commit()
	return &users, http.StatusCreated, "Registered successfully"
}

/* Adding Seeker Profile Repository */
func (r *repository) AddNewSeekerProfile(userId uint, input *dto.InputUserSeekerRegister) (*model.SeekerProfile, string) {
	var seekerProfile model.SeekerProfile
	db := r.db.Model(&seekerProfile)

	if db.Debug().Select("*").Where("user_id = ?", userId).Find(&seekerProfile).RowsAffected > 0 {
		return &seekerProfile, "REGISTER_CONFLICT_409"
	}

	seekerProfile.UserID = userId
	seekerProfile.FirstName = input.FirstName
	seekerProfile.SecondName = input.SecondName

	if db.Debug().Create(&seekerProfile).Error != nil {
		return &seekerProfile, "REGISTER_FAILED_403"
	}
	db.Commit()
	return &seekerProfile, "nil"
}