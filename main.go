package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raisa320/API/config"
	"github.com/raisa320/API/controllers"
	"github.com/raisa320/API/services"
)

func handlers() http.Handler {
	handler := mux.NewRouter()

	handler.HandleFunc("/items", controllers.GetItems).Methods("GET")
	handler.HandleFunc("/items", controllers.SaveItem).Methods("POST")
	return handler
}

func main() {
	port := flag.String("port", ":9000", "The server port")
	flag.Parse()

	services.InitDB()
	services.Db.PingOrDie()

	router := handlers()

	server := config.NewServer(*port, router)

	log.Fatal(server.Start())
}
