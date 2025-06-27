package services

import (
	"e-novel/model"
	"e-novel/model/dto"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type BookService interface {
	GetCategories()([]*model.Category, error)
	GetRecentBook() ([]*dto.BookResponseUser, error)
	GetBookMostLike() ([]*dto.BookResponseUser, error)
	GetBookByCategoryId(id int) ([]*dto.BookResponseUser, error)
	GetBookDetailById(id int)(*dto.ResponseBookDetail, error)

	AddBook(ctx *fiber.Ctx, book *dto.RequestBook, file *multipart.FileHeader, userId int)(int, error)
	AddBookCategory(book *dto.RequestBookCategory) error
}	
