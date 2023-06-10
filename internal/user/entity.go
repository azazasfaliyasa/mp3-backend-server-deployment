package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string
	Username string
	Password string
	RoleID   int
}
