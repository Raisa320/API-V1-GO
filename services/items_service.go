package services

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

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

func GetItem(itemId int) (*models.Item, error) {

	query := `
    SELECT id, customer_name, order_date, product, quantity, price, details
        FROM items WHERE id = $1;
    `
	var item models.Item

	err := Db.QueryRow(query, itemId).Scan(&item.ID, &item.Customer_name, &item.Order_date, &item.Product, &item.Quantity, &item.Price, &item.Details)
	if err != nil {
		if err == sql.ErrNoRows {
			// No se encontró ningún objeto
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func SearchItemByCustomer(customerName string) ([]models.Item, error) {

	query := `
    SELECT id, customer_name, order_date, product, quantity, price, details
        FROM items WHERE LOWER(customer_name) = $1;
    `
	rows, err := Db.Query(query, customerName)
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

	// crear un validador para la estructura Items
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

func changeValue(s interface{}, fieldName string, newValue interface{}) {
	v := reflect.ValueOf(s).Elem()
	f := v.FieldByName(fieldName)

	if f.IsValid() && f.CanSet() && fieldName != "ID" {
		f.Set(reflect.ValueOf(newValue))
	} else {
		fmt.Println("Field not found or cannot be set")
	}
}

func updateDataItem(item *models.Item, dataRequest interface{}) {
	v := reflect.ValueOf(dataRequest)
	t := reflect.TypeOf(dataRequest)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if !value.IsZero() {
			changeValue(item, field.Name, value.Interface())
		}
	}
}

func UpdateItem(ctx context.Context, item *models.Item, dataRequest models.Item) (*models.Item, error) {
	updateDataItem(item, dataRequest)
	query := `UPDATE items
	SET customer_name=$1, order_date=$2, product=$3, quantity=$4, price=$5, details=$6
	WHERE id = $7 ;`

	row := Db.QueryRowContext(
		ctx, query, item.Customer_name, item.Order_date, item.Product, item.Quantity, item.Price, item.Details, item.ID)

	err := row.Err()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return item, nil
}

func DeleteItem(ctx context.Context, itemId int) error {

	query := `
    DELETE FROM items WHERE id = $1;
    `
	row := Db.QueryRowContext(ctx, query, itemId)

	err := row.Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
