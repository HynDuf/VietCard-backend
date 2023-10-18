package card

import (
	"vietcard-backend/internal/delivery/http/request"
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

func (cu *cardUsecase) CreateCard(card *entity.Card) error {
	return cu.cardRepository.CreateCard(card)
}

func (cu *cardUsecase) GetCardByID(id *string) (*entity.Card, error) {
	return cu.cardRepository.GetCardByID(id)
}

func (cu *cardUsecase) UpdateCard(cardID *string, req *request.UpdateCardRequest) (*entity.Card, error) {
	return cu.cardRepository.UpdateCard(cardID, req)
}

func (uc *cardUsecase) UpdateCardReview(card *entity.Card, correct bool) error {
	card.UpdateScheduleSM2(correct)
	return uc.cardRepository.UpdateCardReview(card)
}
