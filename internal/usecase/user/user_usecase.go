package user

import (
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/internal/domain/interface/repository"
	"vietcard-backend/internal/domain/interface/usecase"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) usecase.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}
func (uu *userUsecase) GetUserByEmail(email *string) (*entity.User, error) {
	return uu.userRepository.GetByEmail(email)
}

func (uu *userUsecase) GetUserByID(id *string) (*entity.User, error) {
	return uu.userRepository.GetByID(id)
}
