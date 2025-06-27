package seeders

import (
	"e-novel/model"

	"gorm.io/gorm"
)

func SeedLike(db *gorm.DB) error {
	data := []model.Like{
		{
			UserId: 1,
			BookId: 3,
		},
		{
			UserId: 1,
			BookId: 3,
		},
		{
			UserId: 1,
			BookId: 3,
		},
		{
			UserId: 1,
			BookId: 1,
		},
		{
			UserId: 1,
			BookId: 1,
		},
		{
			UserId: 1,
			BookId: 2,
		},
	}


	for _, v := range data {
		err := db.Table("likes").Create(&v).Error
		if err != nil {
			return err
		}
	}

	return nil
}