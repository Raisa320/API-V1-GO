package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/raisa320/API/api/models"
	"github.com/raisa320/API/api/services"
)

type ItemHandler struct {
	ItemService services.ItemService
}

func (handler *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(handler.ItemService.GetItems())
}

func (handler *ItemHandler) SaveItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = handler.ItemService.SaveItem(ctx, &item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(item)
}
