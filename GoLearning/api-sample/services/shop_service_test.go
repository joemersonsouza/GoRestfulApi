package services

import (
	"golearning/api-sample/models"
	"golearning/api-sample/repository"
	"testing"

	"github.com/jinzhu/gorm"
)

type MockRepo struct{}

var productTest repository.Product = repository.Product{Name: "Test"}
var request = models.ProductRequest{
	Name:     "Test",
	Price:    12,
	Discount: 10,
}

func TestGetItemsShouldReturnItems(t *testing.T) {
	service := NewShopService(newMockRepo())
	products := service.GetItems()

	if len(products) == 0 {
		t.Errorf("Products not fount %d", len(products))
		return
	}
}

func TestCreateItem(t *testing.T) {
	service := NewShopService(newMockRepo())
	service.AddItem(request)
}

func (repo *MockRepo) CreateRepositoryInstance(config *repository.RepositoryConfig) *gorm.DB {
	return nil
}

func (repo *MockRepo) GetProducts() []repository.Product {
	var products []repository.Product
	return append(products, productTest)
}

func (repo *MockRepo) CreateProduct(product repository.Product) {
}

func (repo *MockRepo) UpdateProduct(product repository.Product) {
}

func (repo *MockRepo) GetProductById(id int) repository.Product {
	return productTest
}

func newMockRepo() repository.IShopRepository {
	return new(MockRepo)
}
