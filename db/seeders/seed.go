package seeders

import (

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	SeedBook(db)
	PageSeeder(db)
}