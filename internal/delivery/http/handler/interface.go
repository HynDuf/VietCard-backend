package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type RestHandler interface {
	SignUp(c *gin.Context)
	LogIn(c *gin.Context)
	RefreshToken(c *gin.Context)
	CreateCard(c *gin.Context)
	CreateDeck(c *gin.Context)
	GetDeckWithReviewCards(c *gin.Context)
	UpdateUser(c *gin.Context)
	UpdateCard(c *gin.Context)
	UpdateDeck(c *gin.Context)
	UpdateReviewCards(c *gin.Context)
	CopyDeck(c *gin.Context)
	CopyCardToDeck(c *gin.Context)
	LogInGetAllData(c *gin.Context)
}

func GetLoggedInUserID(c *gin.Context) (string, error) {
	uID, isExisted := c.Get("x-user-id")
	if !isExisted {
		return "", errors.New("Missing x-user-id (set at middleware) in Gin Context")
	}

	return uID.(string), nil
}
