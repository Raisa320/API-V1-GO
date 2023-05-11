package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/raisa320/API/models"
	"github.com/raisa320/API/services"
)

var ErrInvalidInput = errors.New("entrada inválida")

func valiId(id string) (int, error) {
	idInt, error := strconv.Atoi(id)
	if error != nil {
		return 0, ErrInvalidInput
	}
	return idInt, nil
}

func getId(r *http.Request) (int, error) {
	param := mux.Vars(r)
	itemId, err := valiId(param["id"])
	if err != nil {
		return 0, err
	}
	return itemId, nil
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := services.GetItems()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}
	json.NewEncoder(w).Encode(items)
}

func SaveItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)

	//item.Price = math.Round(float64(item.Price)*100) / 100

	_, err := services.SaveItem(r.Context(), &item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itemId, err := getId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	item, err := services.GetItem(itemId)

	if item == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Item not found")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	json.NewEncoder(w).Encode(item)
}

func SearchItemByCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	paramName := r.URL.Query()
	customer, isPresent := paramName["customer"]
	if !isPresent {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Customer Parameter not present")
		return
	}

	item, err := services.SearchItemByCustomer(strings.ToLower(customer[0]))

	if item == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Customer Name not found")
		return
	}

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itemId, err := getId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	item, err := services.GetItem(itemId)

	if item == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Item not found")
		return
	}

	var dataRequest models.Item
	json.NewDecoder(r.Body).Decode(&dataRequest)

	//item.Price = math.Round(float64(item.Price)*100) / 100

	itemUpdated, err := services.UpdateItem(r.Context(), item, dataRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	json.NewEncoder(w).Encode(itemUpdated)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itemId, err := getId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	item, err := services.GetItem(itemId)

	if item == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Item not found")
		return
	}

	err = services.DeleteItem(r.Context(), itemId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	json.NewEncoder(w).Encode("Deleted Successfully")
}
