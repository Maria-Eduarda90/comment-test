package userservices

import (
	"api/internal/dto"
	"api/internal/repository/userepository"
	"api/internal/response"
	"context"
)

func NewUserService(repository userepository.UserRepository) UserService {
	return &service{
		repository,
	}
}

type service struct{
	repository userepository.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, u dto.CreateUserDto) error
	UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error
	GetUserByID(ctx context.Context, id string) (*response.UserResponse, error)
	DeleteUser(ctx context.Context, id string) error
	FindManyUsers(ctx context.Context) (*response.ManyUsersResponse, error)
}