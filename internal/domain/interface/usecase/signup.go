package usecase

import "vietcard-backend/internal/domain/entity"

type SignupUsecase interface {
	Create(user *entity.User) (string, error)
	GetUserByEmail(email *string) (*entity.User, error)
	CreateAccessToken(user *entity.User, secret *string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *entity.User, secret *string, expiry int) (refreshToken string, err error)
}
