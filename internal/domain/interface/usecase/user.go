package usecase

import (
	"vietcard-backend/internal/delivery/http/request"
	"vietcard-backend/internal/domain/entity"
)

type UserUsecase interface {
	GetUserByID(id *string) (*entity.User, error)
	GetUserByEmail(email *string) (*entity.User, error)
	UpdateUser(userID *string, req *request.UpdateUserRequest) (*entity.User, error)
	AddXPToUser(userID *string, XP int) (*entity.User, error)
	CreateFact(*entity.Fact) error
}
