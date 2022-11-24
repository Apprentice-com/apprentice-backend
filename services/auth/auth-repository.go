package authService

import (
	"fmt"

	model "github.com/KadirbekSharau/apprentice-backend/models"
	util "github.com/KadirbekSharau/apprentice-backend/util"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

/* User Login Repository Service */
func (r *repository) UserLogin(input *model.Users) (*model.Users, string) {

	var users model.Users
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Email = input.Email
	users.Password = input.Password

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return &users, <-errorCode
	} else if !users.IsActive {
		errorCode <- "LOGIN_NOT_ACTIVE_403"
		return &users, <-errorCode
	}

	comparePassword := util.ComparePassword(users.Password, input.Password)

	if comparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}

/* Active User Seeker Registration Repository */
func (r *repository) ActiveUserSeekerRegisterRepository(input *InputUserSeekerRegister) (*model.Users, string) {

	var users model.Users
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)
	fmt.Println("Inn!!")

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserAccount.RowsAffected > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
		return &users, <-errorCode
	}
	fmt.Println("Checked!!")

	users.Email = input.Email
	users.Password = input.Password
	users.IsActive = true
	users.UserType = 1
	fmt.Println("Assigned!!")

	addNewUser := db.Debug().Create(&users)
	fmt.Println("Created!!")
	db.Commit()
	fmt.Println("Committed!!")

	if addNewUser.Error != nil {
		errorCode <- "REGISTER_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	r.AddNewSeekerProfile(users.ID, input)

	return &users, <-errorCode
}

/* Active User Employer Registration Repository */
func (r *repository) ActiveUserEmployerRegisterRepository(input *InputUserSeekerRegister) (*model.Users, string) {
	var users model.Users
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)
	fmt.Println("Inn!!")

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)
	fmt.Println("Checked!")

	if checkUserAccount.RowsAffected > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
		return &users, <-errorCode
	}

	// Need to create a employee profile
	users.Email = input.Email
	users.Password = input.Password
	users.IsActive = true
	users.UserType = 2
	fmt.Println("Assigned!")

	addNewUser := db.Debug().Create(&users)
	fmt.Println("Created")
	db.Commit()
	fmt.Println("Comitted")

	if addNewUser.Error != nil {
		errorCode <- "REGISTER_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	// userId, ok := addNewUser.Get("id")
	// if !ok {
	// 	errorCode <- "REGISTER_FAILED_403"
	// 	return &users, <-errorCode
	// }
	// id := userId.(uint)
	// r.AddNewSeekerProfile(id, input)

	return &users, <-errorCode
}

/* Adding Seeker Profile Repository */
func (r *repository) AddNewSeekerProfile(userId uint, input *InputUserSeekerRegister) (*model.SeekerProfiles, string) {

	var seekerProfile model.SeekerProfiles
	db := r.db.Model(&seekerProfile)
	errorCode := make(chan string, 1)

	checkEntity := db.Debug().Select("*").Where("user_id = ?", userId).Find(&seekerProfile)

	if checkEntity.RowsAffected > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
		return &seekerProfile, <-errorCode
	}

	seekerProfile.UserID = userId
	seekerProfile.FirstName = input.FirstName
	seekerProfile.SecondName = input.SecondName

	addNewEntity := db.Debug().Create(&seekerProfile)
	db.Commit()

	if addNewEntity.Error != nil {
		errorCode <- "REGISTER_FAILED_403"
		return &seekerProfile, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &seekerProfile, <-errorCode
}
