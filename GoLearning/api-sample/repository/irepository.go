package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type IRepository interface {
	CreateRepositoryInstance(config *RepositoryConfig) *gorm.DB
	Get() []Product
	GetById(id int) Product
	Update(product Product)
	Create(product Product)
}
