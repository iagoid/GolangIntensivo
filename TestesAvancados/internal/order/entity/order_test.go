package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateANewOrder_ThenShouldReceiveAnErro(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnErro(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnErro(t *testing.T) {
	order := Order{ID: "123", Price: 123}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAValidParams_WhenICallANewOrder_ThenShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order := Order{
		ID:    "123",
		Price: 10.0,
		Tax:   2.0,
	}
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.IsValid())
}

func TestGivenAValidParams_WhenICallANewOrderFunc_ThenShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, err)
}

func TestGivenAPriceAnTax_WhenICallCalculatePrice_ThenShouldSetFinalPrice(t *testing.T) {
	order := Order{
		ID:    "123",
		Price: 10.0,
		Tax:   2.0,
	}
	assert.Nil(t, order.IsValid())
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, 12.0, order.FinalPrice)
}
