package dto

type InputLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type InputUserSeekerRegister struct {
	FirstName  string `json:"name"`
	SecondName string `json:"surname"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,gte=8"`
}