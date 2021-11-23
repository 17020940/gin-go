package models

type User struct {
	ID          *uint   `json:"id" gorm:"primaryKey"`
	Username    *string `json:"username" binding:"required"`
	Password    *string `json:"password" binding:"required"`
	PhoneNumber *string `json:"phoneNumber" binding:"required"`
	Email       *string `json:"email" binding:"required,email"`
}
