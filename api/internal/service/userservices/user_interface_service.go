package userservices

import "api/internal/repository/userepository"

func NewUserService(repository userepository.UserRepository) UserService {
	return &service{
		repository,
	}
}

type service struct{
	repository userepository.UserRepository
}

type UserService interface {
	CreateUser() error
}