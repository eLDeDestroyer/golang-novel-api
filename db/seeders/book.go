package seeders

import (
	"e-novel/model"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"gorm.io/gorm"
)

func SeedBook(db *gorm.DB) {
	books := []model.Book{
		{
			Title: "book 1",
			Description: "",
			ImagePath: "",
			UserId: 1,
		},
		{
			Title: "book 2",
			Description: "",
			ImagePath: "",
			UserId: 1,
		},
	}


	for _,book := range books {
		err := db.Table("books").Create(&book).Error
		if err != nil {
			spew.Dump(err)
		}
	}

	fmt.Println("Seeder Book")
}