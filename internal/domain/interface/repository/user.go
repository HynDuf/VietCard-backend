package repository

import (
	"vietcard-backend/internal/domain/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	GetByEmail(email *string) (*entity.User, error)
	GetByID(id *string) (*entity.User, error)
}
