package seeders

import (
	"e-novel/model"

	"gorm.io/gorm"
)

func SeedBookCategory(db *gorm.DB) error {
	data := []model.BookCategory{
		{
			BookId: 1,
			CategoryId: 1,
		},
		{
			BookId: 1,
			CategoryId: 2,
		},
		{
			BookId: 2,
			CategoryId: 2,
		},
		{
			BookId: 2,
			CategoryId: 3,
		},
		{
			BookId: 3,
			CategoryId: 4,
		},
		{
			BookId: 3,
			CategoryId: 5,
		},
	}

	err := db.Table("book_category").Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}