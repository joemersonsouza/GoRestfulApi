package services

import (
	"golearning/api-sample/models"
	"golearning/api-sample/repository"

	"errors"
	"strconv"
)

type IShopService interface {
	GetItems() []repository.Product
	GetItem(itemId string) (repository.Product, error)
	AddItem(item models.ProductRequest) models.ProductRequest
	UpdateItem(id string, changedItem models.ProductRequest) (repository.Product, error)
}

type ShopService struct {
	Repository repository.IShopRepository
}

func NewShopService(repo repository.IShopRepository) IShopService {
	service := new(ShopService)
	service.Repository = repo

	return service
}

func (shopService *ShopService) GetItems() []repository.Product {
	return shopService.Repository.GetProducts()
}

func (shopService *ShopService) GetItem(itemId string) (repository.Product, error) {
	var id, err = strconv.Atoi(itemId)

	if err != nil {
		return repository.Product{}, err
	}

	var item = shopService.Repository.GetProductById(id)

	if (item != repository.Product{}) {
		return item, nil
	}

	return repository.Product{}, errors.New("Item Not Found")
}

func (shopService *ShopService) AddItem(item models.ProductRequest) models.ProductRequest {

	var product = repository.Product{
		Name:     item.Name,
		Price:    item.Price,
		Discount: item.Discount,
	}

	shopService.Repository.CreateProduct(product)

	return item
}

func (shopService *ShopService) UpdateItem(id string, changedItem models.ProductRequest) (repository.Product, error) {
	var item, err = shopService.GetItem(id)

	if err != nil {
		return repository.Product{}, err
	}

	item.Name = changedItem.Name
	item.Price = changedItem.Price
	item.Discount = changedItem.Discount

	shopService.Repository.UpdateProduct(item)

	return item, nil
}
