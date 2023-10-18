package usecase

import "vietcard-backend/internal/domain/entity"


type RefreshTokenUsecase interface {
	CreateAccessToken(user *entity.User, secret *string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *entity.User, secret *string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken *string, secret *string) (string, error)
}
