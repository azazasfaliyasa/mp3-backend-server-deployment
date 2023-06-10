package user

import (
	"gorm.io/gorm"
	"mini-project-3-azaz/pkg/password"
)

type Repository struct {
	DB *gorm.DB
}

type RepositoryInterface interface {
	ShowAll() ([]User, error)
	Create(user *User) error
	Show(ID string) (*User, error)
	Update(ID string, body User) (*User, error)
	Destroy(ID string) (*User, error)
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) ShowAll() ([]User, error) {
	var users []User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *Repository) Create(user *User) error {
	hash, _ := password.HashPassword(user.Password)
	user.Password = string(hash)
	user.RoleID = 2

	return r.DB.Create(user).Error
}

func (r *Repository) Show(ID string) (*User, error) {
	var user User
	err := r.DB.First(&user, ID).Error
	return &user, err
}

func (r *Repository) Update(ID string, body User) (*User, error) {
	var user User
	err := r.DB.First(&user, ID).Error
	user.FullName = body.FullName
	user.Username = body.Username
	if user.Password != "" {
		hash, _ := password.HashPassword(body.Password)
		user.Password = string(hash)
	}

	query := r.DB.Save(&user).Error
	if query != nil {
		return nil, query
	}

	return &user, err
}

func (r *Repository) Destroy(ID string) (*User, error) {
	var user User
	err := r.DB.First(&user, ID).Error

	query := r.DB.Delete(&user).Error
	if query != nil {
		return nil, query
	}

	return &user, err
}
