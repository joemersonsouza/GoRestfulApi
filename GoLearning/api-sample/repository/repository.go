package repository

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type RepositoryConfig struct {
	Dialect  string
	Host     string
	DBport   string
	User     string
	Database string
	Password string
}

type IShopRepository interface {
	CreateRepositoryInstance(config *RepositoryConfig) *gorm.DB
	GetProducts() []Product
	GetProductById(id int) Product
	UpdateProduct(product Product)
	CreateProduct(product Product)
}

type ShopRepository struct{}

func NewShopRepository() IShopRepository {
	return new(ShopRepository)
}

var once sync.Once

var (
	instance *gorm.DB
)

func (repo *ShopRepository) CreateRepositoryInstance(config *RepositoryConfig) *gorm.DB {

	once.Do(func() {

		instance = _initializeDatabase(config)

	})

	return instance
}

func _initializeDatabase(config *RepositoryConfig) *gorm.DB {
	var err error
	var db *gorm.DB

	dialect := config.Dialect
	host := config.Host
	dbport := config.DBport
	user := config.User
	database := config.Database
	password := config.Password

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, database, password, dbport)

	db, err = gorm.Open(dialect, dbURI)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Println("Connection completed")

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}
