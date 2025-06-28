package seeders

import (
	"e-novel/model"
	"strconv"

	"gorm.io/gorm"
)

func SeedCategory(db *gorm.DB) error {
	for i := 0; i < 5; i++ {
		data := model.Category{
			Category: "category " + strconv.Itoa(i),
		}

		err := db.Table("categories").Create(&data).Error
		if err != nil {
			return err
		}
	}

	return nil
}