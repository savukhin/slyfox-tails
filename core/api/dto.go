package api

type LoginUserDTO struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

type RegisterUserDTO struct {
	Username       string `validate:"required,min=3,max=32" json:"username"`
	Email          string `validate:"required,email,min=6,max=32" json:"email"`
	Password       string `validate:"required,min=3,max=40" json:"password"`
	PasswordRepeat string `validate:"required,min=3,max=40" json:"password_repeat"`
}