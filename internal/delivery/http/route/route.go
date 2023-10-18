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
	_ "vietcard-backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			VietCard Backend API
//	@version		1.0
//	@description	Backend server for VietCard application
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.email	hynduf@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used
func Setup(db *mongo.Database, gin *gin.Engine) {
	userRP := userrepo.NewUserRepository(db)

	loginUsecase := login.NewLoginUsecase(userRP)
	signUpUsecase := signup.NewSignupUsecase(userRP)
	refreshTokenUsecase := refreshtkn.NewRefreshTokenUsecase(userRP)

	h := handler.NewHandler(loginUsecase, signUpUsecase, refreshTokenUsecase)

	publicRouter := gin.Group("")

	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	publicRouter.GET("/swagger/*any", swaggerHandler)
	publicRouter.POST("/api/signup", h.SignUp)
	publicRouter.POST("/api/login", h.LogIn)
	publicRouter.POST("/api/refresh", h.RefreshToken)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(bootstrap.E.AccessTokenSecret))
}
