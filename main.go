package main

import (
	"aggregator_info/handler"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var M0 = "helloworld"

func main() {

	listenPort := 9011

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/handler_1inch", handler.Handler_1inch)
	e.POST("/handler_bancor", handler.HandlerBancor)
	e.POST("/handler_paraswap", handler.HandlerParaswap)
	e.POST("/handler_kyberswap", handler.HandlerKyberswap)
	e.POST("/handler_zeroX", handler.ZeroX_handler)
	e.POST("/handler_mooniswap", handler.Mooniswap_handler)
	e.POST("/handler_dforce", handler.Dforce_handler)
	e.POST("/handler_uniswap_v2", handler.Uniswap_v2_handler)
	e.POST("/handler_sushiswap", handler.Sushiswap_handler) // TODO: ERROR
	// uniswap v1

	// Curve
	// 0x
	// balancer
	// DDEX
	// Loopring
	// paraSwap
	// dYdX
	// DoDo
	// Oasis
	// IDEX
	// DEX.AG
	// Tokenlon

	// listenPort := viper.GetUint32("port")
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", listenPort)); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	// go getData()

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
