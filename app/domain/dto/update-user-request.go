package dto

type UpdateUserDTO struct {
	Name     string  `json:"name" binding:"required"`
	Age      *uint   `json:"age"`
	Address  *string `json:"address"`
	Birthday *string `json:"birthday"` // You can use string for easier date handling (e.g., "YYYY-MM-DD")
}
