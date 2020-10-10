package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/y2labs-0sh/aggregator_info/daemons"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/handler"
	_ "github.com/y2labs-0sh/aggregator_info/logger"
	"github.com/y2labs-0sh/aggregator_info/types"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	viper.SetDefault("port", ":9011")
	viper.SetDefault("tokenslist", "1inch")
}

func main() {
	e := echo.New()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			e.Logger.Fatalf("Fatal error config file: %s \n", err)
		}
	}

	daemonCtx, daemonCancel := context.WithCancel(context.Background())

	uniswapV2PoolDaemon := daemons.NewUniswapV2PoolsDaemon(e.Logger, 200)
	uniswapV2PoolDaemon.Run(daemonCtx)
	tokenPriceDaemon := daemons.NewTokenPriceBalancer(e.Logger)
	tokenPriceDaemon.Run(daemonCtx)
	balancerPoolDaemon := daemons.NewBalancerPoolsDaemon(e.Logger, 200)
	balancerPoolDaemon.Run(daemonCtx)
	mergedPoolDaemon := daemons.NewMergedPoolInfosDaemon(e.Logger)
	mergedPoolDaemon.Run(daemonCtx)

	// Middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &types.EchoContext{Context: c}
			return next(cc)
		}
	})
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/aggrInfo", handler.AggrInfo)
	e.POST("/swapInfo", handler.SwapInfo)
	e.GET("/tokenlist", handler.TokenList)
	e.GET("/tokenicons", handler.TokenIconsList)

	investGroup := e.Group("/invest")
	investHandler := handler.InvestHandler{}
	investGroup.GET("/list", investHandler.Pools)
	investGroup.POST("/prepare", investHandler.Prepare)
	investGroup.POST("/estimate", investHandler.Estimate)
	investGroup.POST("/estimate_prepare", investHandler.EstimateAndPrepare)

	stakingGroup := e.Group("/staking")
	stakingHandler := handler.StakingHandler{}
	stakingGroup.GET("/pools", stakingHandler.Pools)
	stakingGroup.POST("/stake", stakingHandler.Stake)
	stakingGroup.POST("/withdraw", stakingHandler.Withdraw)
	stakingGroup.POST("/claim_reward", stakingHandler.ClaimReward)
	stakingGroup.POST("/exit", stakingHandler.Exit)

	data.GetTokenList(viper.GetString("tokenslist"))

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				log.Info("Updating txFee...")
				estimatetxfee.UpdateTxFee()
				log.Info("Update txFee finished.")
				time.Sleep(3600 * time.Second)
			}
		}
	}(daemonCtx)

	go func() {
		if err := e.Start(viper.GetString("port")); err != nil {
			e.Logger.Fatal(err)
		}
		daemonCancel()
	}()

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
