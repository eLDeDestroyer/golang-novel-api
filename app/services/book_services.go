package services

import (
	"e-novel/model/dto"
)

type BookService interface {
	GetRecentBook()([]*dto.BookResponseUser, error)
	GetBookMostLike()([]*dto.BookResponseUser, error)
}