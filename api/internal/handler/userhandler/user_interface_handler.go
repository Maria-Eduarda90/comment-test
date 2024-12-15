package userhandler

import (
	"api/internal/service/userservices"
	"net/http"
)

func NewUserHandler(service userservices.UserService) UserHandler {
	return &handler{
		service,
	}
} 

type handler struct {
	service userservices.UserService
}

type UserHandler interface{
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	FindManyUsers(w http.ResponseWriter, r *http.Request)
	UpdateUserPassword(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}