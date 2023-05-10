package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/raisa320/API/config"
)

type DbConnection struct {
	*sql.DB //SeudoHerencia
}

var Db DbConnection

func InitDB() {
	err := connect_BD()
	if err != nil {
		log.Fatal(err)
	}
}

// PingOrDie envía un ping a la base de datos y si no se puede alcanzar, registra un error fatal.
func (db *DbConnection) PingOrDie() {
	if err := db.Ping(); err != nil {
		log.Fatalf("no se puede alcanzar la base de datos, error: %v", err)
	}
}

var dbConn *sql.DB

// Connect_BD conecta con la base de datos y devuelve un error si falla la conexión.
func connect_BD() error {

	var errDb error
	dbConfig, errDb := config.LoadEnvVariables()

	if errDb != nil {
		log.Fatalf("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DbName)
	var err error
	dbConn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión exitosa a la base de datos:", dbConn)
	Db = DbConnection{dbConn}
	Db.PingOrDie()
	return err
}
