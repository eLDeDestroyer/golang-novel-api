package repositories

import "e-novel/model/dto"


type UserRepository interface {
	GetMeRepository(id int) ([]map[string]interface{}, error)
	GetUserRepository(username string) ([]map[string]interface{}, error)
	GetPageCountByIdBook(id int)(int, error)

	GetBookLikeUser(userId int)([]map[string]interface{}, error)
	GetBookSaveUser(userId int)([]map[string]interface{}, error)
	GetBookHistoriesUser(userId int)([]map[string]interface{}, error)

	AddBookLikeUser(action *dto.RequestAction) error
	AddBookSaveUser(action *dto.RequestAction) error
	AddBookHistoriesUser(action *dto.RequestAction) error

	CheckBookLikeUser(action *dto.RequestAction) (bool, error)
	CheckBookSaveUser(action *dto.RequestAction) (bool, error)	

	DeleteBookLikeUser(action *dto.RequestAction) error
	DeleteBookSaveUser(action *dto.RequestAction) error
}

