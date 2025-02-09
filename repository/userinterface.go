package repository

import "RestApi-Golang/model"

type UserInterface interface {
	CreateUser(model.User) (string, error)
	GetUserByID(string) (model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUserAgeByID(string, model.User) (int, error)
	DeleteUserByID(string) (int, error)
	DeleteAllUsers() (int, error)
}
