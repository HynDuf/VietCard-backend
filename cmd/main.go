package main

import (
	"vietcard-backend/bootstrap"
	"vietcard-backend/database/mongodb"
	"vietcard-backend/internal/delivery/http/route"

	"github.com/gin-gonic/gin"
)

func main() {
	bootstrap.NewEnv()
	router := gin.Default()

	db := mongodb.NewDBConnection(bootstrap.E.MongoDBURI)
	route.Setup(db, router)
	router.Run(bootstrap.E.ServerAddress)
}
