package app

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"

)


type ProductInterface interface {
	IsValid() (bool, error)
	Enabled() error
	Disabled() error
	GetID() string
	GetName() string
	GetPrice() float64
	GetStatus() string
}

const (
	ENABLED  = "enabled"
	DISABLED = "disabled"
)

type Product struct {
	ID     string `json:"id" valid:"uuidv4"`
	Name   string  `json:"name" valid:"required"`
	Price  float64 `json:"price" valid:"float,optional"`
	Status string  `json:"status" valid:"required"`
}

func NewProduct() *Product {
	product := &Product{
		ID: uuid.NewV4().String(),
		Status: DISABLED,
	}
	return product
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		return false, errors.New("status cannot be empty")
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("status must be either disabled or enabled")
	}

	if p.Price < 0 {
		return false, errors.New("price cannot be less than zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Enabled() error {
	if p.Price >  0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("price must be greater than zero")
}

func (p *Product) Disabled() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("price must be zero to disable the product")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetStatus() string {
	return p.Status
}

