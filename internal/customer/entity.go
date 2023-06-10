package customer

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Avatar    string
}
