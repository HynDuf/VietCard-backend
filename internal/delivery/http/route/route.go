package route

import (
	"vietcard-backend/bootstrap"
	"vietcard-backend/internal/delivery/http/handler"
	"vietcard-backend/internal/delivery/http/middleware"
	"vietcard-backend/internal/repository/cardrepo"
	"vietcard-backend/internal/repository/deckrepo"
	"vietcard-backend/internal/repository/userrepo"
	"vietcard-backend/internal/usecase/card"
	"vietcard-backend/internal/usecase/deck"
	"vietcard-backend/internal/usecase/login"
	"vietcard-backend/internal/usecase/refreshtkn"
	"vietcard-backend/internal/usecase/signup"
	"vietcard-backend/internal/usecase/user"

	_ "vietcard-backend/docs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

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

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Description for what is this security definition being used
func Setup(db *mongo.Database, gin *gin.Engine) {
	userRP := userrepo.NewUserRepository(db)
	cardRP := cardrepo.NewCardRepository(db)
	deckRP := deckrepo.NewDeckRepository(db)

	loginUsecase := login.NewLoginUsecase(userRP)
	signUpUsecase := signup.NewSignupUsecase(userRP)
	refreshTokenUsecase := refreshtkn.NewRefreshTokenUsecase(userRP)
	userUsecase := user.NewUserUsecase(userRP)
	cardUsecase := card.NewCardUsecase(cardRP, deckRP)
	deckUsecase := deck.NewDeckUsecase(deckRP, cardRP, userRP)

	h := handler.NewHandler(loginUsecase, signUpUsecase, refreshTokenUsecase, cardUsecase, deckUsecase, userUsecase)

	publicRouter := gin.Group("")

	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	publicRouter.GET("/swagger/*any", swaggerHandler)
	publicRouter.POST("/api/signup", h.SignUp)
	publicRouter.POST("/api/login", h.LogIn)
	publicRouter.POST("/api/login-get-all", h.LogInGetAllData)
	publicRouter.POST("/api/signup-get-all", h.SignUpGetAllData)
	publicRouter.POST("/api/refresh", h.RefreshToken)
	publicRouter.POST("/api/get-all", h.GetAllData)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(bootstrap.E.AccessTokenSecret))
	protectedRouter.PUT("/api/user/update", h.UpdateUser)
	protectedRouter.POST("/api/card/create", h.CreateCard)
	protectedRouter.PUT("/api/card/update", h.UpdateCard)
	protectedRouter.PUT("/api/card/review", h.UpdateReviewCards)
	protectedRouter.POST("/api/card/copy", h.CopyCardToDeck)
	protectedRouter.POST("/api/deck/create", h.CreateDeck)
	protectedRouter.PUT("/api/deck/update", h.UpdateDeck)
	protectedRouter.DELETE("/api/deck/delete", h.DeleteDeck)
	protectedRouter.GET("/api/deck/review-cards", h.GetDeckWithReviewCards)
	protectedRouter.POST("/api/deck/copy", h.CopyDeck)
	protectedRouter.POST("/api/fact/create", h.CreateFact)
}
