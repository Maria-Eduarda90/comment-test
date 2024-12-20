package userepository

import (
	"api/internal/database/sqlc"
	"api/internal/entity"
	"context"
)

func (r *repository) CreateUser(ctx context.Context, u *entity.UserEntity) error {
	err := r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		ID: u.ID,
		Name: u.Name,
		Email: u.Email,
		Password: u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	user, err := r.queries.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	userEntity := entity.UserEntity{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}
	return &userEntity, nil
}

func (r *repository) FindUserByID(ctx context.Context, id string) (*entity.UserEntity, error) {
	return nil, nil
}

func (r *repository) UpdateUser(ctx context.Context, u *entity.UserEntity) error {
	return nil
}

func (r *repository) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (r *repository) FindManyUsers(ctx context.Context) ([]entity.UserEntity, error) {
	return nil, nil
}

func (r *repository) UpdatePassword(ctx context.Context, pass, id string) error {
	return nil
}

func (r *repository) GetUserPassword(ctx context.Context, id string) (*entity.UserEntity, error) {
	userMock := entity.UserEntity{
		ID: "1",
		Password: "$2a$10$i.f/A6qcKIIldTnGt3ZM8eTgWn4WaK.7MEFoC3WiUtaEuYSiHvoae",
	}

	return &userMock, nil
}