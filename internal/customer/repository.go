package customer

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

type RepositoryInterface interface {
	ShowAll() ([]Customer, error)
	Create(body *Customer) error
	Show(ID string) (*Customer, error)
	Update(ID string, body Customer) (*Customer, error)
	Destroy(ID string) (*Customer, error)
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) ShowAll() ([]Customer, error) {
	var customers []Customer
	err := r.DB.Find(&customers).Error
	return customers, err
}

func (r *Repository) Create(customer *Customer) error {
	return r.DB.Create(customer).Error
}

func (r *Repository) Show(ID string) (*Customer, error) {
	var customer Customer
	err := r.DB.First(&customer, ID).Error
	return &customer, err
}

func (r *Repository) Update(ID string, body Customer) (*Customer, error) {
	var customer Customer
	err := r.DB.First(&customer, ID).Error
	customer.FirstName = body.FirstName
	customer.LastName = body.LastName
	customer.Email = body.Email
	customer.Avatar = body.Avatar

	query := r.DB.Save(&customer).Error
	if query != nil {
		return nil, query
	}

	return &customer, err
}

func (r *Repository) Destroy(ID string) (*Customer, error) {
	var customer Customer
	err := r.DB.First(&customer, ID).Error

	query := r.DB.Delete(&customer).Error
	if query != nil {
		return nil, query
	}

	return &customer, err
}
