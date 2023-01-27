package users

import (
	"context"
	"encoding/json"
	"errors"
	adapter "fiberapi/internal/infraestructure/Adapter"
	"fiberapi/internal/users/domain"
)

type UsersRepository struct {
	adapter adapter.Adapter
}

func NewUsersRepository(s adapter.Adapter) domain.UsersRepository {
	return &UsersRepository{
		adapter: s,
	}
}

func (u UsersRepository) NewPassword(ctx context.Context, password, email string) error {

	//u.adapter.UpdateOne(ctx,)
	update := make(map[string]interface{})

	update["password"] = password

	filter := make(map[string]interface{})

	filter["email"] = email

	err := u.adapter.UpdateOne(ctx, filter, update)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (u UsersRepository) GetEmail(ctx context.Context, email string) (domain.Users, error) {
	var user domain.Users

	results, err := u.adapter.FindEmail(ctx, email)

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
