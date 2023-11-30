package signup

import (
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/internal/domain/interface/repository"
	"vietcard-backend/internal/domain/interface/usecase"
	"vietcard-backend/pkg/tokenutil"
)

type signupUsecase struct {
	userRepository repository.UserRepository
}

func NewSignupUsecase(userRepository repository.UserRepository) usecase.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
	}
}

func (su *signupUsecase) Create(user *entity.User) (string, error) {
	return su.userRepository.Create(user)
}

func (su *signupUsecase) GetUserByEmail(email *string) (*entity.User, error) {
	return su.userRepository.GetByEmail(email)
}

func (su *signupUsecase) CreateAccessToken(user *entity.User, secret *string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *entity.User, secret *string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
