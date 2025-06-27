package seeders

import (
	"e-novel/model"
	"fmt"
	"gorm.io/gorm"
)

func PageSeeder(db *gorm.DB) error {
	data := []model.Page{
		{
			Page: 1,
			Text: "text book 1 page 1",
			BookId: 1,
		},
		{
			Page: 2,
			Text: "text book 1 page 2",
			BookId: 1,
		},
		{
			Page: 3,
			Text: "text book 1 page 2",
			BookId: 1,
		},
		{
			Page: 1,
			Text: "text book 2 page 1",
			BookId: 2,
		},
		{
			Page: 2,
			Text: "text book 2 page 2",
			BookId: 2,
		},
		{
			Page: 1,
			Text: "text book 3 page 1",
			BookId: 3,
		},
	}


	for _, row := range data {
		err := db.Table("pages").Create(&row).Error
		if err != nil {
			return err
		}
	}

	fmt.Println("suces")

	return nil
} 