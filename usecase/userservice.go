package usecase

import (
	"RestApi-Golang/repository"
	"net/http"
)

type UserService struct {
	DBClient repository.UserInterface
}

func (srv UserService) CreateUser(w http.ResponseWriter, r *http.Request) {

}
func (srv UserService) GetUserByID(w http.ResponseWriter, r *http.Request) {

}
func (srv UserService) GetAllUsers(w http.ResponseWriter, r *http.Request) {

}
func (srv UserService) UpdateUserAgeByID(w http.ResponseWriter, r *http.Request) {

}
func (srv UserService) DeleteUserByID(w http.ResponseWriter, r *http.Request) {

}
func (srv UserService) DeleteAllUsers(w http.ResponseWriter, r *http.Request) {

}
