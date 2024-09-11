package dto

type LoginUserDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ForgetPassword struct {
	Email string `json:"email" binding:"required,email"`
}
