package handler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	SignUp(c *gin.Context)
	LogIn(c *gin.Context)
    RefreshToken(c *gin.Context)
}
