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
}