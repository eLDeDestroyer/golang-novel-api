package services

import (
	"e-novel/model"
	"e-novel/model/dto"
)

type PageService interface {
	GetPageBook(bookId int, page int) (*dto.ResponsePage, error)
	AddPageBook(page *model.Page) error
	UpdatePageBook(page *model.Page) error
}