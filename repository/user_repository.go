package repository

import "go-api-server/model"

// UserRepository is interface
type UserRepository interface {
	Insert(user model.User) (id int, err error)
	SelectByID(id int) (model.User, error)
	SelectAll() ([]model.User, error)
}
