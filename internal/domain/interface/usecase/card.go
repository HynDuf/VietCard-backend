package usecase

import "vietcard-backend/internal/domain/entity"

type CardUsecase interface {
	CreateCard(card *entity.Card) error
	GetCardByID(id *string) (*entity.Card, error)
	UpdateCard(card *entity.Card) error
	UpdateCardReview(card *entity.Card, correct bool) error
}
