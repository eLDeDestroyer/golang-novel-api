package repositories

import (
	"e-novel/model"
	"e-novel/model/dto"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (repo *UserRepositoryImpl) GetMeRepository(id int) ([]map[string]interface{}, error) {
	var user []map[string]interface{}

	err := repo.db.Table("users").
		Joins("LEFT JOIN books on books.user_id = users.id").
		Select("users.username as username, users.name as name, books.id as id, books.title as title, books.image_path as image_path, books.description as description").
		Where("users.id = ?", id).
		Find(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (repo *UserRepositoryImpl) GetUserRepository(username string) ([]map[string]interface{}, error) {
	var user []map[string]interface{}

	err := repo.db.Table("users").
		Joins("LEFT JOIN books on books.user_id = users.id").
		Select("users.username as username, users.name as name, books.id as id, books.title as title, books.image_path as image_path, books.description as description").
		Where("users.username = ?", username).
		Find(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) GetPageCountByIdBook(id int) (int, error) {
	var pageCount int64

	err := repo.db.Table("pages").Where("book_id = ?", id).Count(&pageCount).Error

	if err != nil {
		return 0, err
	}

	return int(pageCount), nil
}

func (repo *UserRepositoryImpl) GetBookLikeUser(userId int) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	err := repo.db.Table("books").
		Joins("LEFT JOIN likes on likes.book_id = books.id").
		Select("books.id, books.title, books.description, books.image_path").
		Where("likes.user_id = ?", userId).
		Find(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *UserRepositoryImpl) GetBookSaveUser(userId int) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	err := repo.db.Table("books").
		Joins("LEFT JOIN save_book on save_book.book_id = books.id").
		Select("books.id, books.title, books.description, books.image_path").
		Where("save_book.user_id = ?", userId).
		Find(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *UserRepositoryImpl) GetBookHistoriesUser(userId int) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	err := repo.db.Table("books").
		Joins("LEFT JOIN histories on histories.book_id = books.id").
		Select("books.id, books.title, books.description, books.image_path").
		Where("histories.user_id = ?", userId).
		Find(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *UserRepositoryImpl) AddBookLikeUser(action *dto.RequestAction) error {
	err := repo.db.Table("likes").Create(&action).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepositoryImpl) AddBookSaveUser(action *dto.RequestAction) error {
	err := repo.db.Table("save_book").Create(&action).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepositoryImpl) AddBookHistoriesUser(action *dto.RequestAction) error {
	err := repo.db.Table("histories").Create(&action).Error

	if err != nil {
		return err
	}

	return nil
}



func (repo *UserRepositoryImpl) CheckBookLikeUser(action *dto.RequestAction) (bool, error) {
	var dataCont int64
	var data bool

	err := repo.db.Table("likes").Where("book_id = ? AND user_id = ?", action.BookId, action.BookId).Count(&dataCont).Error

	if err != nil {
		return false, err
	}

	if dataCont != 0 {
		data = true
	} else {
		data = false
	}

	return data, nil
}



func (repo *UserRepositoryImpl) CheckBookSaveUser(action *dto.RequestAction) (bool, error) {
	var dataCont int64
	var data bool

	err := repo.db.Table("save_book").Where("book_id = ? AND user_id = ?", action.BookId, action.BookId).Count(&dataCont).Error

	if err != nil {
		return false, err
	}

	if dataCont != 0 {
		data = true
	} else {
		data = false
	}

	return data, nil
}


func (repo *UserRepositoryImpl) DeleteBookLikeUser(action *dto.RequestAction) error {
	err := repo.db.Table("likes").Where("book_id = ? AND user_id = ?", action.BookId, action.BookId).Delete(&model.Action{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepositoryImpl) DeleteBookSaveUser(action *dto.RequestAction) error {
	err := repo.db.Table("save_book").Where("book_id = ? AND user_id = ?", action.BookId, action.BookId).Delete(&model.Action{}).Error

	if err != nil {
		return err
	}

	return nil
}