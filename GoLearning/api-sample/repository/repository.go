package repository

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type IShopRepository interface {
	CreateRepositoryInstance() *gorm.DB
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

func (repo *ShopRepository) CreateRepositoryInstance() *gorm.DB {

	once.Do(func() {

		instance = _initializeDatabase()

	})

	return instance
}

func _initializeDatabase() *gorm.DB {
	var err error
	var db *gorm.DB

	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbport := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	database := os.Getenv("DATABASE")
	password := os.Getenv("PASSWORD")

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
