package main

import (
	"fmt"
	"log"
	"os"

	database "bank-app/config/database"
	"bank-app/core/routes"
	handlers "bank-app/delivery/http/handlers"
	"bank-app/infrastructure/repositories"
	"bank-app/usecase"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := database.ConnectDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	customerRepo := repositories.NewCustomerRepo(db)
	customerUseCase := usecase.NewCustomerUseCase(customerRepo)
	customerHandler := handlers.NewCustomerHandler(customerUseCase)

	accountRepo := repositories.NewAccountRepo(db)
	accountUseCase := usecase.NewAccountUseCase(accountRepo)
	accountHandler := handlers.NewAccountHandler(accountUseCase, customerUseCase)

	userRepo := repositories.NewUserRepo(db)
	authUseCase := usecase.NewAuthUseCase(userRepo)
	authHandler := handlers.NewAuthHandler(authUseCase)

	r := routes.SetupRouter(customerHandler, accountHandler, authHandler)
	r.Run(":8080")

}
