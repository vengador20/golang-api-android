package domain

import "context"

type UsersRepository interface {
	NewPassword(ctx context.Context, password, email string) error
	GetEmail(ctx context.Context, email string) (Users, error)
}
