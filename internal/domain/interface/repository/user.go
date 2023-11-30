package repository

import (
	"vietcard-backend/internal/delivery/http/request"
	"vietcard-backend/internal/domain/entity"
)

type UserRepository interface {
	Create(user *entity.User) (string, error)
	GetByEmail(email *string) (*entity.User, error)
	GetByID(id *string) (*entity.User, error)
	UpdateUser(userID *string, req *request.UpdateUserRequest) (*entity.User, error)
	UpdateUserXP(user *entity.User) error
}
