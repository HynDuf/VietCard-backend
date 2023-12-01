package user

import (
	"vietcard-backend/internal/delivery/http/request"
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

func (uu *userUsecase) UpdateUser(userID *string, req *request.UpdateUserRequest) (*entity.User, error) {
	return uu.userRepository.UpdateUser(userID, req)
}

func (uu *userUsecase) AddXPToUser(userID *string, XP int) (*entity.User, error) {
	user, err := uu.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	user.XP += XP
	user.UpdateLevel()
	user.UpdateStreak()
	err = uu.userRepository.UpdateUserXP(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu *userUsecase) CreateFact(fact *entity.Fact) error {
	return uu.userRepository.CreateFact(fact)
}
