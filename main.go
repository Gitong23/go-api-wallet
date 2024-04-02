package main

import (
	_ "github.com/KKGo-Software-engineering/fun-exercise-api/docs"
	"github.com/KKGo-Software-engineering/fun-exercise-api/postgres"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			Wallet API
// @version		1.0
// @description	Sophisticated Wallet API
// @host			localhost:1323
func main() {
	p, err := postgres.New()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Add Logger Middleware
	e.Use(middleware.Logger())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	handler := wallet.New(p)

	g := e.Group("/api/v1")
	g.GET("/wallets", handler.WalletHandler)
	g.GET("/users/:id/wallets", handler.WalletUserHandler)
	g.POST("/wallets", handler.CreateWalletHandler)
	g.PUT("/wallets", handler.UpdateWalletHandler)
	g.DELETE("/users/:id/wallets", handler.DeleteWalletHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
