package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
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

func GetItemsPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query() // page - itemsPerPage
	pageParam := params["page"]
	itemsPerPageParam := params["itemsPerPage"]
	if len(pageParam) == 0 {
		pageParam = append(pageParam, "1")
	}
	if len(itemsPerPageParam) == 0 {
		itemsPerPageParam = append(itemsPerPageParam, "5")
	}

	pageNumber, errorPage := strconv.Atoi(pageParam[0])
	itemsPerPage, errorNroItem := strconv.Atoi(itemsPerPageParam[0])
	if errorPage != nil || errorNroItem != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Query params no numbers")
		return
	}

	items, totalItems, err := services.GetItemsPage(pageNumber, itemsPerPage)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(itemsPerPage)))

	// :o Mapa de info de la paginacion
	paginationInfo := map[string]interface{}{
		"totalPages":  totalPages,
		"currentPage": pageNumber,
	}

	// Mapa de respuesta
	responseData := map[string]interface{}{
		"items":      items,
		"pagination": paginationInfo,
	}

	json.NewEncoder(w).Encode(responseData)
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
