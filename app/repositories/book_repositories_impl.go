package repositories

import (
	"e-novel/model"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepositoryImpl(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{
		db: db,
	}
}

func (repo *BookRepositoryImpl) GetCategories() ([]*model.Category, error) {
	var data []*model.Category

	err := repo.db.Table("categories").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *BookRepositoryImpl) GetPageCountBook(id int) (int, error) {
	var pageCount int64

	err := repo.db.Table("pages").Where("book_id = ?", id).Count(&pageCount).Error
	if err != nil {
		return 0, err
	}

	return int(pageCount), nil
}

func (repo *BookRepositoryImpl) GetSeenCountBook(id int) (int, error) {
	var seenCount int64

	err := repo.db.Table("histories").Where("book_id = ?", id).Count(&seenCount).Error
	if err != nil {
		return 0, err
	}

	return int(seenCount), nil
}

func (repo *BookRepositoryImpl) GetLikeCountBook(id int) (int, error) {
	var likeCount int64

	err := repo.db.Table("likes").Where("book_id = ?", id).Count(&likeCount).Error
	if err != nil {
		return 0, err
	}

	return int(likeCount), nil
}

func (repo *BookRepositoryImpl) GetRecentBook() ([]map[string]interface{}, error) {
	var data []map[string]interface{}

	err := repo.db.Table("books").Order("id DESC").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *BookRepositoryImpl) GetBookMostLike() ([]map[string]interface{}, error) {
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

func (repo *BookRepositoryImpl) GetBookByCategoryId(id int) ([]map[string]interface{}, error) {
	var user []map[string]interface{}

	err := repo.db.Table("book_category").
		Joins("LEFT JOIN books on books.id = book_category.book_id").
		Joins("LEFT JOIN pages on pages.book_id = books.id").
		Select("books.id,books.title,books.image_path,books.description,COUNT(pages.book_id) as page_count").
		Group("books.id,books.title,books.image_path,books.description").
		Where("category_id = ?", id).
		Find(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *BookRepositoryImpl) GetBookDetailById(id int) ([]map[string]interface{}, error) {
	var data []map[string]interface{}

	err := repo.db.Table("books").
		Joins("LEFT JOIN pages on pages.book_id = books.id").
		Joins("LEFT JOIN histories on histories.book_id = books.id").
		Joins("LEFT JOIN likes on likes.book_id = books.id").
		Joins("LEFT JOIN users on users.id = books.user_id").
		Select(`books.title,
			   books.description,
			   books.image_path,
			   users.name,
			    pages.id ,
				pages.page, 
				pages.text`).
		Where("books.id = ?", id).
		Find(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *BookRepositoryImpl) AddBook(book *model.Book) (int, error) {
	err := repo.db.Table("books").Create(book).Error

	if err != nil {
		return 0, err
	}

	return int(book.Id), nil
}



func (repo *BookRepositoryImpl) AddBookCategory(book *model.BookCategory) error {
	err := repo.db.Table("book_category").Create(book).Error

	if err != nil {
		return err
	}

	return nil
}
