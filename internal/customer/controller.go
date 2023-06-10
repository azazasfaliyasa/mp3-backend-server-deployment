package customer

import (
	"fmt"
	"gorm.io/gorm"
)

type Controller struct {
	UC UseCaseInterface
}

type ControllerInterface interface {
	ShowAll() (*[]CustomerItem, error)
	Create(body *CreateRequest) (*CustomerItemResponse, error)
	Show(ID string) (*CustomerItem, error)
	Update(ID string, body Customer) (*CustomerItemResponse, error)
	Destroy(ID string) (*CustomerItemResponse, error)
}

func NewController(uc UseCaseInterface) ControllerInterface {
	return &Controller{
		UC: uc,
	}
}

type CustomerItem struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type CustomerItemResponse struct {
	Message string        `json:"message"`
	Data    *CustomerItem `json:"data"`
}

func (c Controller) ShowAll() (*[]CustomerItem, error) {
	customers, err := c.UC.ShowAll()
	if err != nil {
		return nil, err
	}

	res := &[]CustomerItem{}
	for _, customer := range customers {
		c := CustomerItem{
			ID:        customer.ID,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		}
		*res = append(*res, c)
	}

	return res, nil
}

func (c Controller) Create(body *CreateRequest) (*CustomerItemResponse, error) {
	customer := Customer{
		Model:     gorm.Model{},
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Avatar:    body.Avatar,
	}
	err := c.UC.Create(&customer)
	if err != nil {
		return nil, err
	}

	res := &CustomerItemResponse{
		Message: "customer successfully created",
		Data: &CustomerItem{
			ID:        customer.ID,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		},
	}

	return res, nil
}
func (c Controller) Show(ID string) (*CustomerItem, error) {
	customer, err := c.UC.Show(ID)
	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, fmt.Errorf("customer not found")
	}

	res := &CustomerItem{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
	}

	return res, nil
}

func (c Controller) Update(ID string, body Customer) (*CustomerItemResponse, error) {
	customers, err := c.UC.Update(ID, body)
	if err != nil {
		return nil, err
	}

	res := &CustomerItemResponse{
		Message: "customer successfully updated",
		Data: &CustomerItem{
			ID:        customers.ID,
			FirstName: customers.FirstName,
			LastName:  customers.LastName,
			Email:     customers.Email,
			Avatar:    customers.Avatar,
		},
	}

	return res, nil
}

func (c Controller) Destroy(ID string) (*CustomerItemResponse, error) {
	customer, err := c.UC.Destroy(ID)
	if err != nil {
		return nil, err
	}

	res := &CustomerItemResponse{
		Message: "customer successfully destroyed",
		Data: &CustomerItem{
			ID:        customer.ID,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		},
	}

	return res, nil
}
