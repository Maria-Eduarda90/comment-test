package userservices

import (
	"api/internal/dto"
	"api/internal/response"
	"context"
	"time"
)

func (s *service) CreateUser(ctx context.Context, u dto.CreateUserDto) error {
	return nil
}

func (s *service) UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error {
	return nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (*response.UserResponse, error) {
	userFake := response.UserResponse{
		ID: "123",
		Name: "Meyh",
		Email: "teste@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &userFake, nil
}

func (s *service) DeleteUser(ctx context.Context, id string) error {
	return nil
}