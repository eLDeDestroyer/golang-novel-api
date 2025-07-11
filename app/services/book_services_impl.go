package services

import (
	"e-novel/app/repositories"
	"e-novel/model"
	"e-novel/model/dto"
	"fmt"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
)

type BookServiceImpl struct {
	bookRepository repositories.BookRepository
}

func NewBookServiceImpl(bookRepository repositories.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{
		bookRepository: bookRepository,
	}
}

func (service *BookServiceImpl) GetCategories() ([]*model.Category, error) {
	data, err := service.bookRepository.GetCategories()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *BookServiceImpl) GetBookByUsername(title string) ([]*dto.BookResponseUser, error) {
	var dataBook []map[string]interface{}
	var errData error

	if title == "" {
		rawData, err := service.bookRepository.GetAllBook()
		dataBook = rawData
		errData = err
		fmt.Println("text")
	} else {
		rawData, err := service.bookRepository.GetBookByUsername(title)
		dataBook = rawData
		errData = err
	}

	if errData != nil {
		return nil, errData
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

func (service *BookServiceImpl) GetBookByCategoryId(id int) ([]*dto.BookResponseUser, error) {
	dataBook, err := service.bookRepository.GetBookByCategoryId(id)
	if err != nil {
		return nil, err
	}

	spew.Dump(dataBook)
	if len(dataBook) == 0 {
		return nil, fmt.Errorf("fails ")
	}

	datas := []*dto.BookResponseUser{}

	for _, row := range dataBook {
		data := dto.BookResponseUser{
			Id:          int(row["id"].(int64)),
			Title:       row["title"].(string),
			Description: row["description"].(string),
			ImagePath:   row["image_path"].(string),
			PageCount:   int(row["page_count"].(int64)),
		}

		datas = append(datas, &data)
	}

	return datas, nil
}

func (service *BookServiceImpl) GetBookDetailById(id int) (*dto.ResponseBookDetail, error) {
	bookDetail, err := service.bookRepository.GetBookDetailById(id)
	pageCount, err := service.bookRepository.GetPageCountBook(id)
	seenCount, err := service.bookRepository.GetSeenCountBook(id)
	likeCount, err := service.bookRepository.GetLikeCountBook(id)

	if err != nil {
		return nil, err
	}

	if len(bookDetail) == 0 {
		return nil, fmt.Errorf("error le")
	}

	datas := &dto.ResponseBookDetail{
		Name:        bookDetail[0]["name"].(string),
		Title:       bookDetail[0]["title"].(string),
		Description: bookDetail[0]["description"].(string),
		ImagePath:   bookDetail[0]["image_path"].(string),
		Action: dto.ActionResponseBookDetail{
			Seen: seenCount,
			Like: likeCount,
			Page: pageCount,
		},
		Pages: []dto.PagesResponseBookDetail{},
	}

	for _, row := range bookDetail {
		exists := false

		for _, v := range datas.Pages {
			if v.Id == int(row["id"].(int64)) {
				exists = true
				continue
			}
		}

		if exists == false {
			data := dto.PagesResponseBookDetail{
				Id:   int(row["id"].(int64)),
				Page: int(row["page"].(int64)),
				Text: row["text"].(string),
			}

			datas.Pages = append(datas.Pages, data)
		}
	}

	return datas, nil

}

func (service *BookServiceImpl) AddBook(ctx *fiber.Ctx, book *dto.RequestBook, file *multipart.FileHeader, userId int) (int, error) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(99999999)

	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%d%s", randomNumber, ext)
	locationFile := fmt.Sprintf("uploads/%s", newFileName)

	err := ctx.SaveFile(file, fmt.Sprintf("./%s", locationFile))
	if err != nil {
		return 0, err
	}

	data := model.Book{
		Title:       book.Title,
		Description: book.Description,
		ImagePath:   locationFile,
		UserId:      userId,
	}

	bookId, err := service.bookRepository.AddBook(&data)
	if err != nil {
		return 0, err
	}

	return bookId, nil
}

func (service *BookServiceImpl) AddBookCategory(book *dto.RequestBookCategory) error {
	for _, row := range book.CategoryId {
		data := model.BookCategory{
			BookId:     book.BookId,
			CategoryId: row,
		}

		err := service.bookRepository.AddBookCategory(&data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *BookServiceImpl) UpdateBook(ctx *fiber.Ctx, book *dto.RequestBook, file *multipart.FileHeader, userId int, bookId int) (int, error) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(99999999)

	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%d%s", randomNumber, ext)
	locationFile := fmt.Sprintf("uploads/%s", newFileName)

	err := ctx.SaveFile(file, fmt.Sprintf("./%s", locationFile))
	if err != nil {
		return 0, err
	}

	data := model.Book{
		Id:          bookId,
		Title:       book.Title,
		Description: book.Description,
		ImagePath:   locationFile,
		UserId:      userId,
	}

	bookId, err = service.bookRepository.UpdateBook(&data)
	if err != nil {
		return 0, err
	}

	return bookId, nil
}

func (service *BookServiceImpl) UpdateBookCategory(book *dto.RequestBookCategory) error {
	err := service.bookRepository.DeleteBookCategory(book.BookId)
	if err != nil {
		return err
	}

	for _, row := range book.CategoryId {
		data := model.BookCategory{
			BookId:     book.BookId,
			CategoryId: row,
		}

		err := service.bookRepository.AddBookCategory(&data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *BookServiceImpl) DeleteBook(id int) error {
	err := service.bookRepository.DeleteBook(id)
	if err != nil {
		return err
	}

	return nil
}
