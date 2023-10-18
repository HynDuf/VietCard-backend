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
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "User not found"})
		return
	}

	user, err := h.userUsecase.GetUserByID(&id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "User not found"})
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
		UserID:       req.UserID,
		DeckID:       req.DeckID,
		Question:     req.Question,
		Answer:       req.Answer,
		WrongAnswers: req.WrongAnswers,
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
	deck := &entity.Deck{
		UserID: req.UserID,
		Name:   req.Name,
	}
	if user.IsAdmin {
		deck.IsGlobal = true
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
