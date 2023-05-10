package services

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/raisa320/API/models"
)

// GetItems obtiene todos los items de la tabla 'items' de la base de datos.
// Retorna una lista de struct 'models.Item' y un error en caso de que haya ocurrido alguno.
func GetItems() ([]models.Item, error) {
	query := `
    SELECT id, customer_name, order_date, product, quantity, price, details
        FROM items;
    `
	rows, err := Db.Query(query)
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

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return items, nil
}

func SaveItem(ctx context.Context, item *models.Item) (models.Item, error) {

	// crear un validador para la estructura Usuario
	validate := validator.New()

	err := validate.Struct(item)
	if err != nil {
		return *item, err
	}

	query := `INSERT INTO items(
		customer_name, order_date, product, quantity, price, details)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;`

	row := Db.QueryRowContext(
		ctx, query, item.Customer_name, item.Order_date, item.Product, item.Quantity, item.Price, item.Details)

	err = row.Scan(&item.ID)
	if err != nil {
		return *item, err
	}

	return *item, nil
}
