package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"

	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/handler"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	viper.SetDefault("port", ":9011")
	viper.SetDefault("tokens", "uniswap")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	e := echo.New()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		e.Logger.Fatalf("Fatal error config file: %s \n", err)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/aggrInfo", handler.AggrInfo)
	e.POST("/swapInfo", handler.SwapInfo)
	e.GET("/tokenlist", handler.TokenList)

	data.GetTokenList(viper.GetString("tokens"))

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
