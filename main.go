package main

import (
	"aggregator_info/handler"
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

// API Call：
// bancor
// paraswap
// 0x
// tokenlan TODO:change API

// Contract Call：
// Curve
// dforce
// kyber
// mooniswap
// oasis
// 1inch
// Sushiswap
// UniswapV2

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	viper.SetDefault("port", ":9011")
}

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/handler", handler.Handler)

	// dYdX
	// 0x
	// balancer
	// DDEX
	// Loopring
	// DoDo
	// IDEX
	// DEX.AG
	// Tokenlon

	go func() {
		if err := e.Start(viper.GetString("port")); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
