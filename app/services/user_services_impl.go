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

	if userData[0]["id"] != nil {		
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
	}

	if userData[0]["id"] == nil {
		data.Books = nil
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

func (service *UserServiceImpl) GetBookByActionUser(action string, userId int) ([]*dto.BookResponseUser, error) {
	var resData []map[string]interface{}
	var errData error

	if action == "like" {
		userData, err := service.userRepository.GetBookLikeUser(userId)
		errData = err
		resData = userData
	} else if action == "save" {
		userData, err := service.userRepository.GetBookSaveUser(userId)
		errData = err
		resData = userData
	} else {
		userData, err := service.userRepository.GetBookHistoriesUser(userId)
		errData = err
		resData = userData
	}

	if errData != nil {
		return nil, errData
	}

	datas := []*dto.BookResponseUser{}

	if action != "histories" {
		for _, row := range resData {
			bookId := int(row["id"].(int32))
			pageCount, err := service.userRepository.GetPageCountByIdBook(bookId)
			if err != nil {
				return nil, err
			}
	
			data := dto.BookResponseUser{
				Id:          bookId,
				Title:       row["title"].(string),
				Description: row["description"].(string),
				ImagePath:   row["image_path"].(string),
				PageCount:   pageCount,
			}
	
			datas = append(datas, &data)
		}
		
	} else {

		for _, row := range resData {
			bookId := int(row["id"].(int32))
			pageCount, err := service.userRepository.GetPageCountByIdBook(bookId)
			if err != nil {
				return nil, err
			}

			exist := false
			for _, v := range datas {
				if v.Id == bookId {
					exist = true
				}
			}

			if !exist {
				data := dto.BookResponseUser{
					Id:          bookId,
					Title:       row["title"].(string),
					Description: row["description"].(string),
					ImagePath:   row["image_path"].(string),
					PageCount:   pageCount,
				}
		
				datas = append(datas, &data)
			}
		}
	}

	return datas, nil
}


func (service *UserServiceImpl) AddBookActionUser(action string,data *dto.RequestAction) error {
	var resErr error

	if action == "like" {

		check,err := service.userRepository.CheckBookLikeUser(data)
		resErr = err

		if check == false {
			err := service.userRepository.AddBookLikeUser(data)
			resErr = err
		}else {
			err := service.userRepository.DeleteBookLikeUser(data)
			resErr = err
		}

	} else if action == "save" {
		check,err := service.userRepository.CheckBookSaveUser(data)
		resErr = err

		if check == false {
			err := service.userRepository.AddBookSaveUser(data)
			resErr = err
		}else {
			err := service.userRepository.DeleteBookSaveUser(data)
			resErr = err
		}
	} else {
		err := service.userRepository.AddBookHistoriesUser(data)
		resErr = err
	}

	if resErr != nil {
		return resErr
	}

	return nil
}
