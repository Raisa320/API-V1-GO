package repository

import (
	"context"

	"github.com/raisa320/API/api/db"
	"github.com/raisa320/API/api/models"
)

// Repository handle the CRUD operations
type ItemRepository struct {
	Storage *db.Storage
}

var storage = db.New()

func (ir *ItemRepository) GetItems() ([]models.Item, error) {
	query := `
    SELECT id, customer_name, order_date, product, quantity, price, details
        FROM items;
    `
	rows, err := storage.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		rows.Scan(&item.ID, &item.Customer_name, &item.Order_date, &item.Product, &item.Quantity, &item.Price, &item.Details)
		items = append(items, item)
	}

	return items, nil
}

func (ir *ItemRepository) SaveItem(ctx context.Context, item *models.Item) error {
	query := `INSERT INTO items(
		customer_name, order_date, product, quantity, price, details)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;`

	row := storage.DB.QueryRowContext(
		ctx, query, item.Customer_name, item.Order_date, item.Product, item.Quantity, item.Price, item.Details)

	err := row.Scan(&item.ID)
	if err != nil {
		return err
	}

	return nil
}
