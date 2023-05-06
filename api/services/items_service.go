package services

import (
	"context"
	"log"

	"github.com/raisa320/API/api/models"
	"github.com/raisa320/API/api/repository"
)

type ItemService struct {
	ItemRepository repository.ItemRepository
}

func (service *ItemService) GetItems() []models.Item {
	items, error := service.ItemRepository.GetItems()
	if error != nil {
		log.Fatal(error)
	}
	return items
}

func (service *ItemService) SaveItem(ctx context.Context, item *models.Item) error {
	error := service.ItemRepository.SaveItem(ctx, item)
	if error != nil {
		return error
	}
	return nil
}
