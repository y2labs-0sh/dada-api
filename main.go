package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var m1 = make(map[string]Token)
var m2 = make(map[string]Exchange)

func init() {
	constructExchange()
	constructToken()
}

func main() {

	listenPort := 9000

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/handler_1inch", handler_1inch)
	e.POST("/handler_bancor", handlerBancor)
	e.POST("/handler_paraswap", handlerParaswap)
	e.POST("/handler_kyberswap", handlerKyberswap)
	e.POST("/handler_zeroX", zeroX_handler)
	//TODO:  DDEX ， Mooniswap, balancer, uniswap 1, 2, Loopring

	// go func() {
	// 	for {
	// 		getData()
	// 		time.Sleep(3600 * time.Second)
	// 	}
	// }

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
