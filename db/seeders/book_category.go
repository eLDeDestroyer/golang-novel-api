package seeders

import (
	"e-novel/model"
	"fmt"

	"gorm.io/gorm"
)

func SeedBookCategory(db *gorm.DB) error {
	data := []model.BookCategory{
		{
			BookId: 1,
			CategoryId: 1,
		},
		{
			BookId: 2,
			CategoryId: 1,
		},
		{
			BookId: 3,
			CategoryId: 2,
		},
		{
			BookId: 4,
			CategoryId: 3,
		},
		{
			BookId: 5,
			CategoryId: 4,
		},
		{
			BookId: 6,
			CategoryId: 3,
		},
		{
			BookId: 6,
			CategoryId: 4,
		},
		{
			BookId: 7,
			CategoryId: 3,
		},
		{
			BookId: 7,
			CategoryId: 5,
		},
		{
			BookId: 8,
			CategoryId: 6,
		},
	}

	err := db.Table("book_category").Create(&data).Error
	if err != nil {
		return err
	}

	fmt.Println("succes seed book category")

	return nil
}