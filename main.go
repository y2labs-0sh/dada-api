package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"

	"github.com/y2labs-0sh/dada-api/daemons"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	"github.com/y2labs-0sh/dada-api/handler"
	"github.com/y2labs-0sh/dada-api/liquidity_history"
	log "github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/types"
)

var (
	Branch    string
	Commit    string
	BuildTime string
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	viper.SetDefault("port", ":9011")
	viper.SetDefault("tokenslist", "1inch")
}

func main() {
	versionFlag := flag.Bool("version", false, "print the version")
	flag.Parse()
	if *versionFlag {
		fmt.Printf("Branch: %s\n", Branch)
		fmt.Printf("Commit: %s\n", Commit)
		fmt.Printf("BuildTime: %s\n", BuildTime)
		os.Exit(0)
	}

	e := echo.New()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			e.Logger.Fatalf("Fatal error config file: %s \n", err)
		}
	}

	daemonCtx, daemonCancel := context.WithCancel(context.Background())

	tokenListDaemon := daemons.NewTokenListDaemon(e.Logger)
	tokenListDaemon.Run(daemonCtx)
	uniswapV2PoolDaemon := daemons.NewUniswapV2PoolsDaemon(e.Logger, 200)
	uniswapV2PoolDaemon.Run(daemonCtx)
	tokenPriceDaemon := daemons.NewTokenPriceBalancer(e.Logger)
	tokenPriceDaemon.Run(daemonCtx)
	balancerPoolDaemon := daemons.NewBalancerPoolsDaemon(e.Logger, 200)
	balancerPoolDaemon.Run(daemonCtx)
	mergedPoolDaemon := daemons.NewMergedPoolInfosDaemon(e.Logger)
	mergedPoolDaemon.Run(daemonCtx)

	go liquidity_history.Init()

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

	e.GET("/tokenlist", handler.TokenList)
	e.GET("/tokenicons", handler.TokenIconsList)

	e.POST("/aggrInfo", handler.AggrInfo)
	e.POST("/swapInfo", handler.SwapInfo)
	e.POST("/liquidity_history", handler.UniswapLiquidityInfo)
	e.POST("/tx_history", handler.TxHistory)

	investGroup := e.Group("/invest")
	investHandler := handler.InvestHandler{}
	investGroup.GET("/list", investHandler.Pools)
	investGroup.POST("/prepare", investHandler.Prepare)
	investGroup.POST("/estimate", investHandler.Estimate)
	investGroup.POST("/estimate_prepare", investHandler.EstimateAndPrepare)
	investGroup.POST("/multiin", investHandler.MultiAssetsInvest)

	stakingGroup := e.Group("/staking")
	stakingHandler := handler.StakingHandler{}
	stakingGroup.GET("/pools", stakingHandler.Pools)
	stakingGroup.POST("/stake", stakingHandler.Stake)
	stakingGroup.POST("/withdraw", stakingHandler.Withdraw)
	stakingGroup.POST("/claim_reward", stakingHandler.ClaimReward)
	stakingGroup.POST("/exit", stakingHandler.Exit)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				log.Info("Updating txFee...")()
				estimatetxfee.UpdateTxFee()
				log.Info("Update txFee finished")()
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
