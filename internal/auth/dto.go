package auth

type SignUpInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}
