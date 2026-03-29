package main

import (
	"backend/auth"
	"backend/config"
	"backend/middlewares"
	"backend/tickets"
	"backend/users"
	"log"

	_ "backend/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Mini Ticket API
// @version		1.0
// @description	Helpdesk API
// @host			localhost:8080
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Insira o token no formato: Bearer {token}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}

	db := config.ConnectMongo()

	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	userController := users.NewController(userService)

	authService := auth.NewService(userService)
	authController := auth.NewController(authService)

	ticketsRepository := tickets.NewRepository(db)
	ticketsService := tickets.NewService(ticketsRepository)
	ticketsController := tickets.NewController(ticketsService)

	r := gin.Default()
	r.Use(middlewares.ErrorHandler())

	// rotas públicas
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authController.RegisterRoutes(r)
	userController.RegisterPublicRoutes(r)

	// rotas privadas
	private := r.Group("/")
	private.Use(middlewares.AuthHandler())
	userController.RegisterPrivateRoutes(private)
	ticketsController.RegisterRoutes(private)

	r.Run(":8080")
}
