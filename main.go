package main

import (
	"fmt"
	"log"

	"github.com/fauzan264/transaction-api-service/config"
	"github.com/fauzan264/transaction-api-service/handler"
	myMiddleware "github.com/fauzan264/transaction-api-service/middleware"
	"github.com/fauzan264/transaction-api-service/transaction"
	"github.com/fauzan264/transaction-api-service/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadConfig()
	logger, logFile, err := myMiddleware.SetupLogger()
	if err != nil {
		log.Fatalf("error setting up logger: %v", err)
	}
	defer logFile.Close()

	db := config.InitDatabase()
	log.Println(&db)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(myMiddleware.LoggerMiddleware(logger))

	// Repositories
	userRepository := user.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	// Services
	userService := user.NewService(userRepository)
	transactionService := transaction.NewService(transactionRepository, userRepository)

	// Handler
	userHandler := handler.NewUserHandler(userService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	api := e.Group("/api/v1")
	api.POST("/daftar", userHandler.RegisterUser)
	api.GET("/saldo/:number_balance", userHandler.GetBalance)
	api.POST("/tarik", transactionHandler.WithdrawTransaction)
	api.POST("/tabung", transactionHandler.SavingTransaction)

	addr := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)
	e.Logger.Fatal(e.Start(addr))
}