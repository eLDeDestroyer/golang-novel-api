package repositories

import (
	"e-novel/model"
)


type BookRepository interface {
	GetCategories()([]*model.Category, error)
	GetPageCountBook(id int)(int, error)
	GetSeenCountBook(id int)(int, error)
	GetLikeCountBook(id int)(int, error)

	GetRecentBook()([]map[string]interface{}, error)
	GetBookMostLike()([]map[string]interface{}, error)
	GetBookByCategoryId(id int)([]map[string]interface{}, error)
	GetBookDetailById(id int)([]map[string]interface{}, error)
}