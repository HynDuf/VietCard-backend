package usecase

import "vietcard-backend/internal/domain/entity"

type UserUsecase interface {
	GetUserByID(id *string) (*entity.User, error)
	GetUserByEmail(email *string) (*entity.User, error)
}
