package repositories


type BookRepository interface {
	GetPageCountBook(id int)(int, error)
	GetRecentBook()([]map[string]interface{}, error)
	GetBookMostLike()([]map[string]interface{}, error)
}