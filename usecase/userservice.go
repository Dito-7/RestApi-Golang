package usecase

import (
	"RestApi-Golang/dto"
	"RestApi-Golang/model"
	"RestApi-Golang/repository"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserService struct {
	DBClient repository.UserInterface
}

func (srv UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}
	// extract the user data from the request body
	var userReq dto.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)

	if err != nil {
		slog.Error(err.Error())
		res.Error = "Invalid request body"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	user := model.User{
		Name:    userReq.Name,
		Age:     userReq.Age,
		Country: userReq.Country,
	}

	result, err := srv.DBClient.CreateUser(r.Context(), user)
	if err != nil {
		slog.Error(err.Error())
		res.Error = "Failed to create user"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	slog.Info("User created successfully", slog.String("_id", result))
	res.Data = result
	json.NewEncoder(w).Encode(res)
}

func (srv UserService) GetUserByID(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	id := chi.URLParam(r, "id")
	if id == "" {
		slog.Error("ID field is Empty")
		res.Error = "invalid ID"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	user, err := srv.DBClient.GetUserByID(id)
	if err != nil {
		slog.Error(err.Error())
		res.Error = "error while fetching user"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	slog.Info("User fetched successfully")
	res.Data = user

	json.NewEncoder(w).Encode(res)
}

func (srv UserService) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	users, err := srv.DBClient.GetAllUsers()
	if err != nil {
		slog.Error(err.Error())
		res.Error = "error while fetching user"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	slog.Info("User fetched successfully")
	res.Data = users

	json.NewEncoder(w).Encode(res)
}

func (srv UserService) UpdateUserAgeByID(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	id := chi.URLParam(r, "id")
	if id == "" {
		slog.Error("ID field is Empty")
		res.Error = "invalid ID"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	// extract the user data from the request body
	var userReq dto.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)

	if err != nil {
		slog.Error(err.Error())
		res.Error = "Invalid request body"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	result, err := srv.DBClient.UpdateUserAgeByID(id, userReq.Age)
	if err != nil {
		slog.Error(err.Error())
		res.Error = "error while updating user"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	slog.Info("User update successfully")
	res.Data = result

	json.NewEncoder(w).Encode(res)
}

func (srv UserService) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	id := chi.URLParam(r, "id")
	if id == "" {
		slog.Error("ID field is Empty")
		res.Error = "invalid ID"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	result, err := srv.DBClient.DeleteUserByID(id)
	if err != nil {
		slog.Error(err.Error())
		res.Error = "error while deleting user"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	slog.Info("User delete successfully")
	res.Data = result

	json.NewEncoder(w).Encode(res)
}

func (srv UserService) DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	result, err := srv.DBClient.DeleteAllUsers()
	if err != nil {
		slog.Error(err.Error())
		res.Error = "error while deleting user"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	slog.Info("User delete successfully")
	res.Data = result

	json.NewEncoder(w).Encode(res)
}

func New(dbclient repository.UserInterface) UserService {
	return UserService{DBClient: dbclient}
}
