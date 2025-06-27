package services

import (
	"e-novel/model"
	"e-novel/model/dto"
)

type BookService interface {
	GetCategories()([]*model.Category, error)
	GetRecentBook() ([]*dto.BookResponseUser, error)
	GetBookMostLike() ([]*dto.BookResponseUser, error)
	GetBookByCategoryId(id int) ([]*dto.BookResponseUser, error)
	GetBookDetailById(id int)(*dto.ResponseBookDetail, error)
}
