package main

import (
	"fmt"
	"log"

	"github.com/fauzan264/transaction-api-service/config"
	"github.com/fauzan264/transaction-api-service/handler"
	"github.com/fauzan264/transaction-api-service/middleware"
	"github.com/fauzan264/transaction-api-service/user"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()
	logger, logFile, err := middleware.SetupLogger()
	if err != nil {
		log.Fatalf("error setting up logger: %v", err)
	}
	defer logFile.Close()

	db := config.InitDatabase()
	log.Println(&db)

	e := echo.New()
	e.Use(middleware.LoggerMiddleware(logger))

	// Repositories
	userRepository := user.NewRepository(db)

	// Services
	userService := user.NewService(userRepository)

	// Handler
	userHandler := handler.NewAuthHandler(userService)

	api := e.Group("/api/v1")
	// api.POST("/user/register", authHandler.RegisterUser)
	api.POST("/user/daftar", userHandler.RegisterUser)
	
	// /api/daftar // nama, nik, no_hp. 
		// 200 field no_rekening generate yang berisi data nomor rekening nasabah
		// 400 deskripsi kesalahan terkait data yg dikirim
	// /api/tabung // no_rekening, nominal
		// 200 saldo yang berisi data saldo nasabah saat ini
		// 400 remark yang berisi deskripsi kesalahan terkait data yg dikirim
	// /api/tarik // no_rekening, nominal
		// 200 field saldo yang berisi data saldo nasabah saat ini.
		// 400 field remark yang berisi deskripsi kesalahan terkait data yg dikirim
	// /api/saldo/{{no_rekening}}
		// field saldo yang berisi data saldo nasabah saat ini
		// field remark yang berisi deskripsi kesalahan terkait data yg dikirim

	addr := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)
	e.Logger.Fatal(e.Start(addr))
}