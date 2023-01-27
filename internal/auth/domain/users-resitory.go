package domain_users

import "context"

type UserRepository interface {
	GetEmail(ctx context.Context, email string) (User, error)
	CreateUser(ctx context.Context, user UserRegister) error
}
