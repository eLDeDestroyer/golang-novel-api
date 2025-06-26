package services

import (
	"e-novel/app/repositories"
	"e-novel/model/dto"
	"fmt"
)

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserServiceImpl(userRepository repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (service *UserServiceImpl) GetMeService(id int) (*dto.ResponseUser, error) {
	userData, err := service.userRepository.GetMeRepository(id)
	if err != nil {
		return nil, err
	}

	if len(userData) == 0 {
		return nil, fmt.Errorf("fails get data")
	}


	data := dto.ResponseUser{
		Username: userData[0]["username"].(string),
		Name:     userData[0]["name"].(string),
		Books:    []dto.BookResponseUser{},
	}

	for _, row := range userData {
		bookId := int(row["id"].(int64))
		pageCount, err := service.userRepository.GetPageCountByIdBook(bookId)

		if err != nil {
			return nil, err
		}

		book := dto.BookResponseUser{
			Id:          bookId,
			Title:       row["title"].(string),
			ImagePath:   row["image_path"].(string),
			Description: row["description"].(string),
			PageCount:   pageCount,
		}

		data.Books = append(data.Books, book)
	}

	return &data, nil
}

func (service *UserServiceImpl) GetUserService(username string) (*dto.ResponseUser, error) {
	userData, err := service.userRepository.GetUserRepository(username)
	if err != nil {
		return nil, err
	}

	if len(userData) == 0 {
		return nil, fmt.Errorf("fails get data")
	}

	data := dto.ResponseUser{
		Username: userData[0]["username"].(string),
		Name:     userData[0]["name"].(string),
		Books:    []dto.BookResponseUser{},
	}

	for _, row := range userData {
		bookId := int(row["id"].(int64))
		pageCount, err := service.userRepository.GetPageCountByIdBook(bookId)
		if err != nil {
			return nil, err
		}

		book := dto.BookResponseUser{
			Id:          bookId,
			Title:       row["title"].(string),
			ImagePath:   row["image_path"].(string),
			Description: row["description"].(string),
			PageCount:   pageCount,
		}

		data.Books = append(data.Books, book)
	}

	return &data, err
}
