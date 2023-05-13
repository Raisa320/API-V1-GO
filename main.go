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
	handler.HandleFunc("/itemsPerPage", controllers.GetItemsPage).Methods("GET")
	handler.HandleFunc("/items", controllers.SaveItem).Methods("POST")
	handler.HandleFunc("/items/{id}", controllers.GetItem).Methods("GET")
	handler.HandleFunc("/item", controllers.SearchItemByCustomer).Methods("GET")
	handler.HandleFunc("/items/{id}", controllers.UpdateItem).Methods("PUT")
	handler.HandleFunc("/items/{id}", controllers.DeleteItem).Methods("DELETE")
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
