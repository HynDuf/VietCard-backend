package route

import (
	"vietcard-backend/bootstrap"
	"vietcard-backend/internal/delivery/http/handler"
	"vietcard-backend/internal/delivery/http/middleware"
	"vietcard-backend/internal/repository/userrepo"
	"vietcard-backend/internal/usecase/login"
	"vietcard-backend/internal/usecase/refreshtkn"
	"vietcard-backend/internal/usecase/signup"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(db *mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	userRP := userrepo.NewUserRepository(db)
	loginUsecase := login.NewLoginUsecase(userRP)
	signUpUsecase := signup.NewSignupUsecase(userRP)
	refreshTokenUsecase := refreshtkn.NewRefreshTokenUsecase(userRP)
	h := handler.NewHandler(loginUsecase, signUpUsecase, refreshTokenUsecase)
	publicRouter.POST("/api/signup", h.SignUp)
	publicRouter.POST("/api/login", h.LogIn)
	publicRouter.POST("/api/refresh", h.RefreshToken)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(bootstrap.E.AccessTokenSecret))
}
