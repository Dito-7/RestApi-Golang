package repository

import (
	"RestApi-Golang/model"
	"context"
)

type UserInterface interface {
	CreateUser(context.Context, model.User) (string, error)
	GetUserByID(string) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUserAgeByID(string, int) (int, error)
	DeleteUserByID(string) (int, error)
	DeleteAllUsers() (int, error)
}
