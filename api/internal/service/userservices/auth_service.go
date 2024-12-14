package userservices

import (
	"api/internal/dto"
	"api/internal/response"
	"context"
	"errors"
	"log/slog"
)

func (s *service) Login(ctx context.Context, u dto.LoginDTO) (*response.UserAuthToken, error) {
	user, err := s.repository.FindUserByEmail(ctx, u.Email)
	if err != nil {
		slog.Error("error to search user by email", "err", err, slog.String("package", "userservices"))
		return nil, errors.New("error to search user password")
	}

	if user == nil {
		slog.Error("user not found", slog.String("package", "userservices"))
		return nil, errors.New("user not found")
	}

	return nil, nil
}