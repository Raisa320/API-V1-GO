
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
Body Request

```json
{
    "name":"James",
    "orderDate":"2022-01-02T15:04:05Z",
    "product":"Product 3",
    "quantity":22,
    "price":82.96
}

```


## Run Locally

Start the server

```bash
  go run main.go
```
## Authors

- [Raisa Orellana](https://github.com/Raisa320)

