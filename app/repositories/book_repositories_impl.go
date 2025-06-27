package repositories

import "gorm.io/gorm"

type BookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepositoryImpl(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{
		db: db,
	}
}


func (repo *BookRepositoryImpl) GetPageCountBook(id int)(int, error) {
	var pageCount int64

	err := repo.db.Table("pages").Where("book_id = ?", id).Count(&pageCount).Error
	if err != nil {
		return 0, err
	}
	
	return int(pageCount), nil
}


func (repo *BookRepositoryImpl) GetRecentBook()([]map[string]interface{}, error) {
	var data []map[string]interface{}
	
	err := repo.db.Table("books").Order("id DESC").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *BookRepositoryImpl) GetBookMostLike()([]map[string]interface{}, error) {
	var data []map[string]interface{}
	
	err := repo.db.Table("books").
		Joins("LEFT JOIN likes on likes.book_id = books.id").
		Select("books.id, books.title, books.description, books.image_path, COUNT(likes.book_id) as like_count").
		Group("books.id, books.title, books.description, books.image_path").
		Order("like_count DESC").
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}