package repositories

import (
	"e-novel/model"
	"e-novel/model/dto"

	"gorm.io/gorm"
)

type PageRepositoryImpl struct {
	db *gorm.DB
}

func NewPageRepositoryImpl(db *gorm.DB) *PageRepositoryImpl {
	return &PageRepositoryImpl{
		db: db,
	}
}

func (repo *PageRepositoryImpl) GetPageBookCount(bookId int) (int, error) {
	var pageCount int64

	err := repo.db.Table("pages").Where("book_id = ?", bookId).Count(&pageCount).Error
	if err != nil {
		return 0, err
	}

	return int(pageCount), nil
}

func (repo *PageRepositoryImpl) GetPageBook(request *dto.RequestPage) ([]map[string]interface{}, error) {
	var rawData []map[string]interface{}
	
	err := 	repo.db.Table("pages").Where("book_id = ? AND page IN ?", request.BookId, request.Page).Find(&rawData).Error
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func (repo *PageRepositoryImpl) AddPageBook(page *model.Page)  error {
	err := 	repo.db.Table("pages").Create(&page).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *PageRepositoryImpl) UpdatePageBook(page *model.Page)  error {
	err := 	repo.db.Table("pages").Where("book_id = ? AND page = ?", page.BookId, page.Page).Updates(&page).Error
	if err != nil {
		return err
	}

	return nil
}