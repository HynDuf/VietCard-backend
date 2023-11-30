package handler

import (
	"net/http"
	"vietcard-backend/bootstrap"
	"vietcard-backend/internal/delivery/http/request"
	"vietcard-backend/internal/delivery/http/response"
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/internal/domain/interface/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type restHandler struct {
	loginUsecase        usecase.LoginUsecase
	signUpUsecase       usecase.SignupUsecase
	refreshTokenUsecase usecase.RefreshTokenUsecase
	cardUsecase         usecase.CardUsecase
	deckUsecase         usecase.DeckUsecase
	userUsecase         usecase.UserUsecase
}

func NewHandler(loginUc usecase.LoginUsecase, signUpUc usecase.SignupUsecase, refreshTokenUc usecase.RefreshTokenUsecase, cardUc usecase.CardUsecase, deckUc usecase.DeckUsecase, userUc usecase.UserUsecase) RestHandler {
	return &restHandler{
		loginUsecase:        loginUc,
		signUpUsecase:       signUpUc,
		refreshTokenUsecase: refreshTokenUc,
		cardUsecase:         cardUc,
		deckUsecase:         deckUc,
		userUsecase:         userUc,
	}
}

// SignUp	godoc
// SignUp	API
//
//	@Summary		Sign Up
//	@Description	Sign Up
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@Router			/api/signup [post]
//	@Param			signup_request	formData	request.SignupRequest	true	"Sign Up Request"
//	@Success		200				{object}	response.SignupResponse
//	@Failure		400				{object}	response.ErrorResponse
//	@Failure		409				{object}	response.ErrorResponse
//	@Failure		500				{object}	response.ErrorResponse
func (h *restHandler) SignUp(c *gin.Context) {
	var request request.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := h.signUpUsecase.GetUserByEmail(&request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}
	if user != nil {
		c.JSON(http.StatusConflict, response.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user = &entity.User{
		Name:           request.Name,
		Email:          request.Email,
		HashedPassword: request.Password,
	}

	err = h.signUpUsecase.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, er := h.signUpUsecase.CreateAccessToken(user, &bootstrap.E.AccessTokenSecret, bootstrap.E.AccessTokenExpiryHour)
	if er != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := h.signUpUsecase.CreateRefreshToken(user, &bootstrap.E.RefreshTokenSecret, bootstrap.E.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := response.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}

// LogIn	godoc
// LogIn	API
//
//	@Summary		Log In
//	@Description	Log In
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@Router			/api/login [post]
//	@Param			login_request	formData	request.LoginRequest	true	"Log In Request"
//	@Success		200				{object}	response.LoginResponse
//	@Failure		400				{object}	response.ErrorResponse
//	@Failure		401				{object}	response.ErrorResponse
//	@Failure		404				{object}	response.ErrorResponse
//	@Failure		500				{object}	response.ErrorResponse
func (h *restHandler) LogIn(c *gin.Context) {
	var request request.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := h.userUsecase.GetUserByEmail(&request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := h.loginUsecase.CreateAccessToken(user, &bootstrap.E.AccessTokenSecret, bootstrap.E.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := h.loginUsecase.CreateRefreshToken(user, &bootstrap.E.RefreshTokenSecret, bootstrap.E.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}

// RefreshToken	godoc
// RefreshToken	API
//
//	@Summary		Refresh Token
//	@Description	Refresh Token
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Router			/api/refresh [post]
//	@Param			refresh_token_request	body		request.RefreshTokenRequest	true	"Refresh Token Request"
//	@Success		200						{object}	response.RefreshTokenResponse
//	@Failure		401						{object}	response.ErrorResponse
//	@Failure		500						{object}	response.ErrorResponse
func (h *restHandler) RefreshToken(c *gin.Context) {
	var request request.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	id, err := h.refreshTokenUsecase.ExtractIDFromToken(&request.RefreshToken, &bootstrap.E.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := h.userUsecase.GetUserByID(&id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := h.refreshTokenUsecase.CreateAccessToken(user, &bootstrap.E.AccessTokenSecret, bootstrap.E.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := h.refreshTokenUsecase.CreateRefreshToken(user, &bootstrap.E.RefreshTokenSecret, bootstrap.E.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenResponse := response.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}

// CreateCard	godoc
// CreateCard	API
//
//	@Summary		Create New Card
//	@Description	Create New Card
//	@Tags			card
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/card/create [post]
//	@Param			create_card_request	body		request.CreateCardRequest	true	"Create Card Request"
//	@Success		200					{object}	response.CreateCardResponse
//	@Failure		500					{object}	response.ErrorResponse
func (h *restHandler) CreateCard(c *gin.Context) {
	var (
		req request.CreateCardRequest
		err error
	)

	uID, err := GetLoggedInUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}
	req.UserID, err = primitive.ObjectIDFromHex(uID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}
	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}
	if len(req.WrongAnswers) < 3 {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Must have at least 3 wrong answers"})
		return
	}

	card := &entity.Card{
		UserID:           req.UserID,
		DeckID:           req.DeckID,
		QuestionImgURL:   req.QuestionImgURL,
		QuestionImgLabel: req.QuestionImgLabel,
		Question:         req.Question,
		Answer:           req.Answer,
		WrongAnswers:     req.WrongAnswers,
	}
	err = h.cardUsecase.CreateCard(card)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	createCardResponse := response.CreateCardResponse{
		Success: true,
	}

	c.JSON(http.StatusOK, createCardResponse)
}

// CreateDeck	godoc
// CreateDeck	API
//
//	@Summary		Create New Deck
//	@Description	Create New Deck
//	@Tags			deck
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/deck/create [post]
//	@Param			create_deck_request	body		request.CreateDeckRequest	true	"Create Deck Request"
//	@Success		200					{object}	response.CreateDeckResponse
//	@Failure		500					{object}	response.ErrorResponse
func (h *restHandler) CreateDeck(c *gin.Context) {
	var (
		req request.CreateDeckRequest
		err error
	)

	uID, err := GetLoggedInUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}
	req.UserID, err = primitive.ObjectIDFromHex(uID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := h.userUsecase.GetUserByID(&uID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: "User ID doesn't exist in DB"})
		return
	}
	deck := &entity.Deck{
		UserID:              req.UserID,
		Name:                req.Name,
		Description:         req.Description,
		DescriptionImageURL: req.DescriptionImageURL,
		TotalCards:          req.TotalCards,
	}
	if user.IsAdmin {
		deck.IsPublic = true
	}
	err = h.deckUsecase.CreateDeck(deck)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	createDeckResponse := response.CreateDeckResponse{
		Success: true,
	}

	c.JSON(http.StatusOK, createDeckResponse)
}

// GetDeckWithReviewCards	godoc
// GetDeckWithReviewCards	API
//
//	@Summary		Get Deck With Review Cards Of Logged In User
//	@Description	Get Deck With Review Cards Of Logged In User
//	@Tags			deck
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/deck/review-cards [get]
//	@Success		200	{object}	[]entity.DeckWithReviewCards
//	@Failure		500	{object}	response.ErrorResponse
func (h *restHandler) GetDeckWithReviewCards(c *gin.Context) {
	uID, err := GetLoggedInUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	deckWithCards, err := h.deckUsecase.GetReviewCardsAllDecksOfUser(&uID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, deckWithCards)
}

// UpdateUser	godoc
// UpdateUser	API
//
//	@Summary		Update User Details
//	@Description	Update User Details
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/user/update [put]
//	@Param			update_user_request	body		request.UpdateUserRequest	true	"Update User Request"
//	@Success		200					{object}	response.UpdateUserResponse
//	@Failure		400					{object}	response.ErrorResponse
//	@Failure		500					{object}	response.ErrorResponse
func (h *restHandler) UpdateUser(c *gin.Context) {
	var (
		req request.UpdateUserRequest
		err error
	)

	uID, err := GetLoggedInUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}
	if req.HashedPassword != nil {
		encryptedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(*req.HashedPassword),
			bcrypt.DefaultCost,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
			return
		}
		pass := string(encryptedPassword)
		req.HashedPassword = &pass
	}

	user, err := h.userUsecase.UpdateUser(&uID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	resp := response.UpdateUserResponse{
		Success: true,
		User:    *user,
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateCard	godoc
// UpdateCard	API
//
//	@Summary		Update Card Details
//	@Description	Update Card Details
//	@Tags			card
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/card/update [put]
//	@Param			update_card_request	body		request.UpdateCardRequest	true	"Update Card Request"
//	@Success		200					{object}	response.UpdateCardResponse
//	@Failure		400					{object}	response.ErrorResponse
//	@Failure		500					{object}	response.ErrorResponse
func (h *restHandler) UpdateCard(c *gin.Context) {
	var (
		req request.UpdateCardRequest
		err error
	)

	uID, err := GetLoggedInUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	cardID := req.CardID.Hex()
	card, err := h.cardUsecase.GetCardByID(&cardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}
	if card.UserID.Hex() != uID {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Not your card! Can't update! Logged in user != card's user"})
		return
	}

	req.CardID = nil
	card, err = h.cardUsecase.UpdateCard(&cardID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	resp := response.UpdateCardResponse{
		Success: true,
		Card:    *card,
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateDeck	godoc
// UpdateDeck	API
//
//	@Summary		Update Deck Details
//	@Description	Update Deck Details
//	@Tags			deck
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/deck/update [put]
//	@Param			update_deck_request	body		request.UpdateDeckRequest	true	"Update Deck Request"
//	@Success		200					{object}	response.UpdateDeckResponse
//	@Failure		400					{object}	response.ErrorResponse
//	@Failure		500					{object}	response.ErrorResponse
func (h *restHandler) UpdateDeck(c *gin.Context) {
	var (
		req request.UpdateDeckRequest
		err error
	)

	uID, err := GetLoggedInUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	deckID := req.DeckID.Hex()
	deck, err := h.deckUsecase.GetDeckByID(&deckID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}
	if deck.UserID.Hex() != uID {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Not your deck! Can't update! Logged in user != deck's user"})
		return
	}

	req.DeckID = nil
	deck, err = h.deckUsecase.UpdateDeck(&deckID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	resp := response.UpdateDeckResponse{
		Success: true,
		Deck:    *deck,
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateReviewCards	godoc
// UpdateReviewCards	API
//
//	@Summary		Update Review Cards
//	@Description	Update Review Cards
//	@Tags			card
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/card/review [put]
//	@Param			update_review_cards_request	body		request.UpdateReviewCardsRequest	true	"Update Review Cards Request"
//	@Success		200							{object}	response.UpdateReviewCardsResponse
//	@Failure		400							{object}	response.ErrorResponse
//	@Failure		500							{object}	response.ErrorResponse
func (h *restHandler) UpdateReviewCards(c *gin.Context) {
	var (
		req request.UpdateReviewCardsRequest
		err error
	)

	uID, err := GetLoggedInUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	deckID := req.DeckID.Hex()
	deck, err := h.deckUsecase.GetDeckByID(&deckID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}
	if deck.UserID.Hex() != uID {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Not your deck! Can't update! Logged in user != deck's user"})
		return
	}

	if len(req.CardIDs) == 0 || len(req.CardIDs) != len(req.IsCorrect) {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Invalid card_ids[] or is_correct[] parameters"})
		return
	}

	deck.UpdateReview()
	cards, _, _, _, err := h.cardUsecase.GetReviewCardsByDeck(&deckID, deck.MaxNewCards-deck.CurNewCards, deck.MaxReviewCards-deck.CurReviewCards)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}
	needUpdate := make(map[string]bool)
	cardsMap := make(map[string]*entity.Card)
	for _, card := range *cards {
		cardsMap[card.ID.Hex()] = &card
	}
	for i, id := range req.CardIDs {
		correct := req.IsCorrect[i]
		c := cardsMap[id.Hex()]
		c.UpdateScheduleSM2(correct)
		needUpdate[id.Hex()] = true
		if correct {
			if c.NumReviews == 1 {
				deck.CurNewCards++
			} else {
				deck.CurReviewCards++
			}
		}
	}

	for id, update := range needUpdate {
		if update {
			err := h.cardUsecase.UpdateCardReview(cardsMap[id])
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
				return
			}
		}
	}

	cards, numBlueCards, numGreenCards, numRedCards, err := h.cardUsecase.GetReviewCardsByDeck(&deckID, deck.MaxNewCards-deck.CurNewCards, deck.MaxReviewCards-deck.CurReviewCards)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}
	updateDeckReq := &request.UpdateDeckRequest{
		LastReview:  &deck.LastReview,
		CurNewCards: &deck.CurNewCards,
		MaxNewCards: &deck.MaxNewCards,
	}
	deck, err = h.deckUsecase.UpdateDeck(&deckID, updateDeckReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}
	_, err = h.userUsecase.AddXPToUser(&uID, req.TotalXP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	resp := response.UpdateReviewCardsResponse{
		Success:       true,
		Cards:         *cards,
		NumBlueCards:  numBlueCards,
		NumRedCards:   numRedCards,
		NumGreenCards: numGreenCards,
	}
	c.JSON(http.StatusOK, resp)
}

// CopyDeck	godoc
// CopyDeck	API
//
//	@Summary		Copy Deck
//	@Description	Copy Deck
//	@Tags			deck
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/deck/copy [post]
//	@Param			copy_deck_request	body		request.CopyDeckRequest	true	"Copy Deck Request"
//	@Success		200					{object}	response.CopyDeckResponse
//	@Failure		400					{object}	response.ErrorResponse
//	@Failure		500					{object}	response.ErrorResponse
func (h *restHandler) CopyDeck(c *gin.Context) {
	var (
		req request.CopyDeckRequest
		err error
	)

	uID, err := GetLoggedInUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	deckID := req.DeckID.Hex()
	deck, err := h.deckUsecase.GetDeckByID(&deckID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}
	if !deck.IsPublic && deck.UserID.Hex() != uID {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Can't copy private deck! Deck is not global and Logged in user != deck's user"})
		return
	}

	err = h.deckUsecase.CopyDeck(&uID, &deckID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	resp := response.CopyDeckResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, resp)
}

// CopyCardToDeck	godoc
// CopyCardToDeck	API
//
//	@Summary		Copy Card To Deck
//	@Description	Copy Card To Deck
//	@Tags			card
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/card/copy [post]
//	@Param			copy_card_to_deck_request	body		request.CopyCardToDeckRequest	true	"Copy Card To Deck Request"
//	@Success		200							{object}	response.CopyCardToDeckResponse
//	@Failure		400							{object}	response.ErrorResponse
//	@Failure		500							{object}	response.ErrorResponse
func (h *restHandler) CopyCardToDeck(c *gin.Context) {
	var (
		req request.CopyCardToDeckRequest
		err error
	)

	// uID, err := GetLoggedInUserID(c)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	// 	return
	// }

	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	cardID := req.CardID.Hex()
	deckID := req.DeckID.Hex()
	err = h.cardUsecase.CopyCardToDeck(&cardID, &deckID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	resp := response.CopyDeckResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, resp)
}

// LogInGetAllData	godoc
// LogInGetAllData	API
//
//	@Summary		Log In And Get All Data
//	@Description	Log In And Get All Data
//	@Tags			mobile
//	@Accept			multipart/form-data
//	@Produce		json
//	@Router			/api/login-get-all [post]
//	@Param			login_request	formData	request.LoginRequest	true	"Log In Request"
//	@Success		200				{object}	response.LoginGetAllDataResponse
//	@Failure		400				{object}	response.ErrorResponse
//	@Failure		401				{object}	response.ErrorResponse
//	@Failure		404				{object}	response.ErrorResponse
//	@Failure		500				{object}	response.ErrorResponse
func (h *restHandler) LogInGetAllData(c *gin.Context) {
	var request request.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := h.userUsecase.GetUserByEmail(&request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := h.loginUsecase.CreateAccessToken(user, &bootstrap.E.AccessTokenSecret, bootstrap.E.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := h.loginUsecase.CreateRefreshToken(user, &bootstrap.E.RefreshTokenSecret, bootstrap.E.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	uID := user.ID.Hex()
	userDecks, publicDecks, decksWithReviewCard, err := h.deckUsecase.GetDecksWithCards(&uID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := response.LoginGetAllDataResponse{
		AccessToken:          accessToken,
		RefreshToken:         refreshToken,
		User:                 *user,
		UserDeckAndCards:     *userDecks,
		PublicDeckAndCards:   *publicDecks,
		DecksWithReviewCards: *decksWithReviewCard,
	}

	c.JSON(http.StatusOK, loginResponse)
}

// GetAllData	godoc
// GetAllData	API
//
//	@Summary		Get All Data
//	@Description	Get All Data
//	@Tags			mobile
//	@Accept			multipart/form-data
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/get-all [get]
//	@Success		200				{object}	response.GetAllDataResponse
//	@Failure		500				{object}	response.ErrorResponse
func (h *restHandler) GetAllData(c *gin.Context) {
	uID, err := GetLoggedInUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := h.userUsecase.GetUserByID(&uID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	userDecks, publicDecks, decksWithReviewCard, err := h.deckUsecase.GetDecksWithCards(&uID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := response.GetAllDataResponse{
		User:                 *user,
		UserDeckAndCards:     *userDecks,
		PublicDeckAndCards:   *publicDecks,
		DecksWithReviewCards: *decksWithReviewCard,
	}

	c.JSON(http.StatusOK, loginResponse)
}
