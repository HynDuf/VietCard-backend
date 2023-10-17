package repository

import (
	"errors"
	"vietcard-backend/internal/domain/entity"
)

var ErrUserNotFound = errors.New("User not found")

type UserRepository interface {
	Create(user *entity.User) error
	GetByEmail(email *string) (*entity.User, error)
	GetByID(id *string) (*entity.User, error)
}
