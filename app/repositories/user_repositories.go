package repositories


type UserRepository interface {
	GetMeRepository(id int) ([]map[string]interface{}, error)
	GetUserRepository(username string) ([]map[string]interface{}, error)
	GetPageCountByIdBook(id int)(int, error)
}

