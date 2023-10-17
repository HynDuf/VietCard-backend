package login

import (
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/internal/domain/interface/repository"
	"vietcard-backend/internal/domain/interface/usecase"
	"vietcard-backend/pkg/tokenutil"
)

type loginUsecase struct {
	userRepository repository.UserRepository
}

func NewLoginUsecase(userRepository repository.UserRepository) usecase.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
	}
}
func (lu *loginUsecase) GetUserByEmail(email *string) (*entity.User, error) {
	return lu.userRepository.GetByEmail(email)
}

func (lu *loginUsecase) CreateAccessToken(user *entity.User, secret *string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *entity.User, secret *string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
