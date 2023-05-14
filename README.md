
# API GO

Bootcamp de programación en GO, ejercicio de aplicación de creación de un API.

## Tech Stack

**Server:** Go

**Database:** Postgres

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DB_HOST`

`DB_PORT`

`DB_NAME`

`DB_USER`

`DB_PASSWORD`

## API Reference

#### Get all items

```http
  GET /items
```

#### Post item

```http
  POST /items
```
`Body Request`

```json
{
    "name":"James",
    "orderDate":"2022-01-02T15:04:05Z",
    "product":"Product 3",
    "quantity":22,
    "price":82.96
}

```

#### Get One Item
```http
  GET /items/{id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch |

#### Update item
```http
  PUT /items/{id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch |

`Example Body Request`

```json
{
    "name":"James",
    "orderDate":"2022-01-02T15:04:05Z",
    "product":"Product 3",
    "quantity":22,
    "price":82.96
}
```

#### Delete item
```http
  DELETE /items/{id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch |

#### Search item by customer
```http
  GET /items?customer={customer}
```
| Query Params | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `customer`      | `string` | **Required**. Customer name of item to fetch |


#### Get Items Pagination
```http
  GET /itemsPerPage?page={page}&itemsPerPage={itemsPerPage}
```
| Query Params | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `page`      | `int` | **Optional**. Number of page   |
| `itemsPerPage`      | `int` | **Optional**. Total items per page |

## Run Locally

Start the server

```bash
  go run main.go
```
## Authors

- [Raisa Orellana](https://github.com/Raisa320)

