package app_test

import (
	"github.com/stretchr/testify/require"
	"github.com/turnes/hexagonal-architecture/app"
	"testing"
)

func TestProduct_Enabled (t *testing.T) {
	product := app.Product{
		Name: "Product1",
		Price: 2.5,
		Status: app.DISABLED,
	}
	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "price must be greater than zero", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := app.Product{
		Name: "Product1",
		Price: 0,
		Status: app.ENABLED,
	}
	err := product.Disable()
	require.Nil(t, err)

	product.Price = 2.5
	err = product.Disable()
	require.Equal(t, "price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{
		Name: "Product1",
		Price: 10,
		Status: "",
	}

	_, err := product.IsValid()
	require.Equal(t, "status cannot be empty", err.Error())

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "status must be either disabled or enabled", err.Error())

	product.Status = app.ENABLED
	_, err = product.IsValid()
	require.Nil(t,err)

	product.Status = app.DISABLED
	_, err = product.IsValid()
	require.Nil(t,err)

	product.Price = -1
	_, err = product.IsValid()
	require.Equal(t, "price cannot be less than zero", err.Error())
}