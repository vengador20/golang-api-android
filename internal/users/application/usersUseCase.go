package application

import "fiberapi/internal/users/domain"

type UsersCase struct {
	domain.UsersRepository
}

func (u *UsersCase) NewPassword(password string) error {
	return nil
}
