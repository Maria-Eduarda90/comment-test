package userepository

import (
	"api/internal/database/sqlc"
	"api/internal/entity"
	"context"
	"database/sql"
	"time"
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
	user, err := r.queries.FindUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	userEntity := entity.UserEntity{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &userEntity, nil
}

func (r *repository) UpdateUser(ctx context.Context, u *entity.UserEntity) error {
	var name sql.NullString
	var email sql.NullString

	if u.Name != "" {
		name = sql.NullString{String: u.Name, Valid: true}
	} else {
		name = sql.NullString{Valid: false}
	}

	if u.Email != "" {
		email = sql.NullString{String: u.Email, Valid: true}
	} else {
		email = sql.NullString{Valid: false}
	}

	err := r.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:        u.ID,
		Name:      name.String,
		Email:     email.String,
		UpdatedAt: u.UpdatedAt,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteUser(ctx context.Context, id string) error {
	err := r.queries.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FindManyUsers(ctx context.Context) ([]entity.UserEntity, error) {
	users, err := r.queries.FindManyUsers(ctx)
	if err != nil {
		return nil, err
	}

	var usersEntity []entity.UserEntity

	for _, user := range users {
		userEntity := entity.UserEntity {
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		usersEntity = append(usersEntity, userEntity)
	}
	return usersEntity, nil
}

func (r *repository) UpdatePassword(ctx context.Context, pass, id string) error {
	err := r.queries.UpdatePassword(ctx, sqlc.UpdatePasswordParams{
		ID: id,
		Password: pass,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetUserPassword(ctx context.Context, id string) (string, error) {
	pass, err := r.queries.GetUserPassword(ctx, id)

	if err != nil {
		return "", err
	}

	return pass, nil
}