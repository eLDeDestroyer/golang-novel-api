package services

import (
	"e-novel/model"
	"e-novel/model/dto"
)

type AuthService interface {
	LoginService(username string, password string) (*dto.ResponseLogin, error)
	RegisterService(user *model.User) error
}
