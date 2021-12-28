package models

import (
	"github.com/google/uuid"
)

// Order represents a collection of products that should be shipped to the specified shipping address
type Order struct {
	ID       uuid.UUID `json:"id,omitempty" validate:"required,uuid"`
	Products []Product `json:"products" validate:"required,dive"`
	Customer Customer  `json:"customer" validate:"required,dive"`
}

// Product represents a single product in an order
type Product struct {
	ProductCode string `json:"productCode" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
}

// Address represents an customers address
type Address struct {
	Line1      string `json:"line1" validate:"required"`
	Line2      string `json:"line2,omitempty"`
	City       string `json:"city" validate:"required"`
	State      string `json:"state" validate:"required"`
	PostalCode string `json:"postalCode" validate:"required"`
}

// Customer represents information about the customer placing the order
type Customer struct {
	FirstName       string  `json:"firstName" validate:"required"`
	LastName        string  `json:"lastName" validate:"required"`
	EmailAddress    string  `json:"emailAddress" validate:"required"`
	ShippingAddress Address `json:"shippingAddress" validate:"required,dive"`
}
