package seeders

import (

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	SeedUser(db)
	SeedBook(db)
	PageSeeder(db)
	SeedLike(db)
	SeedCategory(db)
	SeedBookCategory(db)
}