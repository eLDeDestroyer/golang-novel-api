package repositories

import "e-novel/model"

type AuthRepository interface {
	LoginRepository(username string) (*model.User, error)
	RegisterRepository(user *model.User) error
}
