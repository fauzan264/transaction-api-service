package main

import (
	"fmt"
	"log"

	"github.com/fauzan264/transaction-api-service/config"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()
	db := config.InitDatabase()
	log.Println(db)

	e := echo.New()
	
	addr := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)
	e.Logger.Fatal(e.Start(addr))
}