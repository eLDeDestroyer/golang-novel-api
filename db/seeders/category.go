package seeders

import (
	"e-novel/model"
	"fmt"

	"gorm.io/gorm"
)

func SeedCategory(db *gorm.DB) error {
	data := []model.Category{
		{
			Category: "romance",
		},
		{
			Category: "drama",
		},
		{
			Category: "adventure",
		},
		{
			Category: "horror",
		},
		{
			Category: "fantasy",
		},
		{
			Category: "comedy",
		},
	}
	for _, v := range data {

		err := db.Table("categories").Create(&v).Error
		if err != nil {
			return err
		}
	}

	fmt.Println("succes seed category")


	return nil
}
