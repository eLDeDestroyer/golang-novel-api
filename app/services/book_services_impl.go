package services

import (
	"e-novel/app/repositories"
	"e-novel/model/dto"
	"fmt"
)

type BookServiceImpl struct {
	bookRepository repositories.BookRepository
}

func NewBookServiceImpl(bookRepository repositories.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{
		bookRepository: bookRepository,
	}
}

func (service *BookServiceImpl) GetRecentBook() ([]*dto.BookResponseUser, error) {
	dataBook, err := service.bookRepository.GetRecentBook()
	if err != nil {
		return nil, err
	}

	if len(dataBook) == 0 {
		return nil, fmt.Errorf("fails ")
	}

	datas := []*dto.BookResponseUser{}

	for _, row := range dataBook {
		bookId := int(row["id"].(int32))
		pageCount, err := service.bookRepository.GetPageCountBook(bookId)
		if err != nil {
			return nil, err
		}

		data := &dto.BookResponseUser{
			Id:          bookId,
			Title:       row["title"].(string),
			Description: row["description"].(string),
			ImagePath:   row["image_path"].(string),
			PageCount:   pageCount,
		}

		datas = append(datas, data)
	}

	return datas, nil
}




func (service *BookServiceImpl) GetBookMostLike() ([]*dto.BookResponseUser, error) {
	dataBook, err := service.bookRepository.GetBookMostLike()
	if err != nil {
		return nil, err
	}

	if len(dataBook) == 0 {
		return nil, fmt.Errorf("fails ")
	}

	datas := []*dto.BookResponseUser{}

	for _, row := range dataBook {
		bookId := int(row["id"].(int32))
		pageCount, err := service.bookRepository.GetPageCountBook(bookId)
		if err != nil {
			return nil, err
		}

		data := &dto.BookResponseUser{
			Id:          bookId,
			Title:       row["title"].(string),
			Description: row["description"].(string),
			ImagePath:   row["image_path"].(string),
			PageCount:   pageCount,
		}

		datas = append(datas, data)
	}

	return datas, nil
}
