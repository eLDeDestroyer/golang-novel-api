package repositories

import (
	"e-novel/model"
	"e-novel/model/dto"
)

type PageRepository interface {
	GetPageBookCount(bookId int) (int, error)
	GetPageBook(request *dto.RequestPage) ([]map[string]interface{}, error)

	AddPageBook(page *model.Page) error
	UpdatePageBook(page *model.Page) error
}
