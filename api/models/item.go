package models

import "time"

type Item struct {
	ID            string    `json:"id"`
	Customer_name string    `json:"name"`
	Order_date    time.Time `json:"orderDate"`
	Product       string    `json:"product"`
	Quantity      int       `json:"quantity"`
	Price         float32   `json:"price"`
	Details       string    `json:"details,omitempty"`
}
