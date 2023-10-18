package card

import (
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/internal/domain/interface/repository"
	"vietcard-backend/internal/domain/interface/usecase"
)

type cardUsecase struct {
	cardRepository repository.CardRepository
}

func NewCardUsecase(cr repository.CardRepository) usecase.CardUsecase {
	return &cardUsecase{
		cardRepository: cr,
	}
}
func (uc *cardUsecase) CreateCard(card *entity.Card) error {
	return uc.cardRepository.CreateCard(card)
}
func (uc *cardUsecase) GetCardByID(id *string) (*entity.Card, error) {
	return uc.cardRepository.GetCardByID(id)
}
func (uc *cardUsecase) UpdateCard(card *entity.Card) error {
	return uc.cardRepository.UpdateCard(card)
}
func (uc *cardUsecase) UpdateCardReview(card *entity.Card, correct bool) error {
	card.UpdateScheduleSM2(correct)
	return uc.cardRepository.UpdateCardReview(card)
}
