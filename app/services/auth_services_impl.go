package services

import (
	"e-novel/app/repositories"
	"e-novel/helper"
	"e-novel/model"
	"e-novel/model/dto"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	authRepository repositories.AuthRepository
}

func (service *AuthServiceImpl) LoginService(username string, password string) (*dto.ResponseLogin, error) {
	var response dto.ResponseLogin

	user, err := service.authRepository.LoginRepository(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	tokenString, err := helper.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	response.Token = tokenString
	response.Username = user.Username
	
	return &response, nil
}

func (service *AuthServiceImpl) RegisterService(user *model.User) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashPassword)
	user.IsPremium = false
	return service.authRepository.RegisterRepository(user)
}
