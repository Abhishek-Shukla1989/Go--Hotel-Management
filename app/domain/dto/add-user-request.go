package dto

type CreateUserDTO struct {
	Name     string  `json:"name" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	Password string  `json:"password" binding:"required"`
	Age      *uint   `json:"age"`
	Address  *string `json:"address"`
	Birthday *string `json:"birthday"` // You can use string for easier date handling (e.g., "YYYY-MM-DD")
}

type UserDTO struct {
	Name     string  `json:"name" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	Age      *uint   `json:"age"`
	Address  *string `json:"address"`
	Birthday *string `json:"birthday"` // You can use string for easier date handling (e.g., "YYYY-MM-DD")
}
