package services

import "e-novel/model/dto"


type UserService interface {
	GetMeService(id int) (*dto.ResponseUser, error)
	GetUserService(username string) (*dto.ResponseUser, error)
}

