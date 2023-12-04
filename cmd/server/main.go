package main

import (
	"fmt"
	"golang_track_expense/application"
	"golang_track_expense/infrastructure/postgres"
	"golang_track_expense/interfaces"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed to load .env file: %v", err)
		panic(fmt.Sprintf("Failed to load .env file: %v", err))
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	db, err := postgres.ConnectPostgres(host, port, user, pass, dbname)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db not connected")
	}

	router := gin.Default()
	userRepository := postgres.NewUserRepository(db)
	userService := application.NewUserService(userRepository)
	userHandlers := interfaces.NewUserHandler(userService)

	transactionRepository := postgres.NewTransactionRepository(db)
	transactionService := application.NewTransactionService(transactionRepository)
	transactionHandlers := interfaces.NewTransactionHandler(transactionService)

	interfaces.SetRoutes(router, userHandlers, transactionHandlers)

	router.Run(":8080")

}
