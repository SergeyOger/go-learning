package models

type User struct {
	ID    uint32 `gorm:"primary_key" json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
}
