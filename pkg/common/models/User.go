package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	IsUser  bool `json:"is_user"`
	IsAdmin bool `json:"is_admin"`
}
