package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func (repo *ShopRepository) getDb() *gorm.DB {
	return repo.CreateRepositoryInstance(nil)
}

func (repo *ShopRepository) CreateProduct(product Product) {
	repo.getDb().Create(&product)
}

func (repo *ShopRepository) UpdateProduct(product Product) {
	repo.getDb().Save(&product)
}

func (repo *ShopRepository) GetProducts() []Product {
	var product []Product
	repo.getDb().Find(&product)

	return product
}

func (repo *ShopRepository) GetProductById(id int) Product {
	var product Product
	repo.getDb().First(&product, id)

	return product
}
