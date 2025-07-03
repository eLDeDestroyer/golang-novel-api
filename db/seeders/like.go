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
			UserId: 2,
			BookId: 3,
		},
		{
			UserId: 3,
			BookId: 3,
		},
		{
			UserId: 4,
			BookId: 3,
		},
		{
			UserId: 5,
			BookId: 3,
		},
		{
			UserId: 6,
			BookId: 3,
		},

		{
			UserId: 1,
			BookId: 2,
		},
		{
			UserId: 2,
			BookId: 2,
		},
		{
			UserId: 3,
			BookId: 2,
		},
		{
			UserId: 4,
			BookId: 2,
		},
		{
			UserId: 5,
			BookId: 2,
		},


		{
			UserId: 1,
			BookId: 5,
		},
		{
			UserId: 2,
			BookId: 5,
		},
		{
			UserId: 3,
			BookId: 5,
		},
		{
			UserId: 4,
			BookId: 5,
		},


		{
			UserId: 1,
			BookId: 7,
		},
		{
			UserId: 2,
			BookId: 7,
		},
		{
			UserId: 3,
			BookId: 7,
		},


		{
			UserId: 1,
			BookId: 8,
		},
		{
			UserId: 2,
			BookId: 8,
		},



		{
			UserId: 1,
			BookId: 1,
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