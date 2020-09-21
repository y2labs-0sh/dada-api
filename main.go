package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"

	"github.com/y2labs-0sh/aggregator_info/daemons"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/handler"
	"github.com/y2labs-0sh/aggregator_info/types"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	viper.SetDefault("port", ":9011")
	viper.SetDefault("tokens", "uniswap")
}

func main() {
	e := echo.New()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		e.Logger.Fatalf("Fatal error config file: %s \n", err)
	}

	daemonCtx, daemonCancel := context.WithCancel(context.Background())

	uniswapListDaemon := daemons.NewUniswapV2Daemon(e.Logger, 100)
	uniswapListDaemon.Run(daemonCtx)

	daemonsMap := make(map[string]daemons.Daemon)
	daemonsMap[daemons.UniswapV2ListDaemonName] = uniswapListDaemon

	// Middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &types.EchoContext{Context: c}
			cc.SetDaemonsMap(daemonsMap)
			return next(cc)
		}
	})
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/aggrInfo", handler.AggrInfo)
	e.POST("/swapInfo", handler.SwapInfo)
	e.GET("/tokenlist", handler.TokenList)

	investGroup := e.Group("/invest")
	investGroup.GET("/list", handler.InvestList)
	investGroup.POST("/prepare", handler.PrepareInvest)

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
	}(daemonCtx)

	go func(ctx context.Context) {
		if err := e.Start(viper.GetString("port")); err != nil {
			e.Logger.Fatal(err)
		}
	}(daemonCtx)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	daemonCancel()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
