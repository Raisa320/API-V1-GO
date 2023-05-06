package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	storage *Storage
	once    sync.Once
)

type Storage struct {
	DB *sql.DB
}

func getConnection() (*sql.DB, error) {
	//uri := os.Getenv("DATABASE_URI")
	return sql.Open("postgres", "postgres://postgres:password@localhost/labora-proyect-1?sslmode=disable")
}

func initDB() {
	db, err := getConnection()
	if err != nil {
		log.Panic(err)
	}

	storage = &Storage{
		DB: db,
	}
}

func New() *Storage {
	once.Do(initDB)
	return storage
}

func (db *Storage) Close() error {
	return db.DB.Close()
}
