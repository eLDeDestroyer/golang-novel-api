package repositories

import (
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
