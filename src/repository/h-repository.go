package repository

import (
	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	model "github.com/KadirbekSharau/apprentice-backend/src/models"
)

/* Active User Employer Registration Repository */
func (r *repository) ActiveUserEmployerRegister(input *dto.InputUserSeekerRegister) (*model.Users, string) {
	var users model.Users
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserAccount.RowsAffected > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
		return &users, <-errorCode
	}

	// Need to create a employee profile
	users.Email = input.Email
	users.Password = input.Password
	users.IsActive = true
	users.UserType = 2

	addNewUser := db.Debug().Create(&users)
	db.Commit()

	if addNewUser.Error != nil {
		errorCode <- "REGISTER_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}