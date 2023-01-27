package auth

import (
	"context"
	"encoding/json"
	"errors"
	domain_users "fiberapi/internal/auth/domain"
	adapter "fiberapi/internal/infraestructure/Adapter"
)

type AuthRepository struct {
	adapter adapter.Adapter
}

func NewAuthRepository(s adapter.Adapter) domain_users.UserRepository {
	return &AuthRepository{
		adapter: s,
	}
}

func (a AuthRepository) GetEmail(ctx context.Context, email string) (domain_users.User, error) {
	var user domain_users.User

	results, err := a.adapter.FindEmail(ctx, email)

	if err != nil {
		return user, errors.New(err.Error())
	}

	res, err := json.Marshal(results)

	if err != nil {
		return user, errors.New(err.Error())
	}

	json.Unmarshal(res, &user)

	return user, nil
}

func (a AuthRepository) CreateUser(ctx context.Context, user domain_users.UserRegister) error {
	return nil
}
