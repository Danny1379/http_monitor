package models

type User struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Username     string `json:"username"`
	PasswordHash string
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
