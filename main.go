package main

import (
	"crowfunding_api/db"
	"crowfunding_api/handler"
	"crowfunding_api/user"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	userRepo := user.NewRepository(db.DB)
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)

	router.Run(":8080")

}
