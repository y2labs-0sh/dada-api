package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"

	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/handler"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	viper.SetDefault("port", ":9011")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/aggrInfo", handler.AggrInfo)
	e.POST("/swapInfo", handler.SwapInfo)

	e.GET("/tokenlist", handler.TokenList)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				estimatetxfee.UpdateTxFee()
				time.Sleep(600 * time.Second)
			}
		}
	}(ctx)

	go func(ctx context.Context) {
		if err := e.Start(viper.GetString("port")); err != nil {
			e.Logger.Fatal(err)
		}
	}(ctx)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
