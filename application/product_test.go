package application_test

import (
	"testing"

	"github.com/elielsimoncelo/fc-hexagonal-arch/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	valid, err := product.IsValid()
	require.Nil(t, err)
	require.Equal(t, true, valid)

	product.Status = "INVALID"
	valid, err = product.IsValid()
	require.Equal(t, "The status must be enabled or disabled", err.Error())
	require.Equal(t, false, valid)

	product.Status = application.ENABLED
	valid, err = product.IsValid()
	require.Nil(t, err)
	require.Equal(t, true, valid)

	product.Price = -50
	valid, err = product.IsValid()
	require.Equal(t, "The price must be greatear or equal zero", err.Error())
	require.Equal(t, false, valid)
}
