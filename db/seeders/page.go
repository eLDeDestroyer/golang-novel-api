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
			BookId: 7,
		},
		{
			Page: 2,
			Text: "text book 1 page 2",
			BookId: 7,
		},
		{
			Page: 3,
			Text: "text book 1 page 2",
			BookId: 7,
		},
		{
			Page: 1,
			Text: "text book 2 page 1",
			BookId: 8,
		},
		{
			Page: 2,
			Text: "text book 2 page 2",
			BookId: 8,
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