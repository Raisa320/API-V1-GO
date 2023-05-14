package models

import "time"

type Item struct {
	ID            string    `json:"id"`
	Customer_name string    `json:"name" validate:"required"`
	Order_date    time.Time `json:"orderDate" validate:"required"`
	Product       string    `json:"product" validate:"required"`
	Quantity      int       `json:"quantity" validate:"required"`
	Price         float64   `json:"price" validate:"required"`
	Details       *string   `json:"details,omitempty"`
	TotalPrice    float64   `json:"totalPrice"`
}
