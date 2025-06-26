package repositories

import (
	"e-novel/model"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}


func NewAuthRepositoryImpl(db *gorm.DB) *AuthRepositoryImpl{
	return &AuthRepositoryImpl{
		db: db,
	}
}

func (repo *AuthRepositoryImpl) LoginRepository(username string) (*model.User, error) {
	var user model.User
	err := repo.db.Table("users").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *AuthRepositoryImpl) RegisterRepository(user *model.User) error {
	err := repo.db.Table("users").Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
