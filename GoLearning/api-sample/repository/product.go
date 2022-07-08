package repository

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model

	Name     string `gorm:"typevarchar(255);inique_index`
	Price    float32
	Discount int
}
