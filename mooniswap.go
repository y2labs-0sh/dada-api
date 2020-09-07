package main

import (
	contractabi "aggregator_info/contract_abi"
	"errors"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"
)

// 已知合约:
// DAI-USDC 0x31631b3dd6c697e574d6b886708cd44f5ccf258f
// ETH-DAI 0x75116bd1ab4b0065b44e1a4ea9b4180a171406ed
// ETH-LEND 0x3863fc8c1cc59f160280f5d3e4c1a4c63f945ce3
// ETH-USDC 0x61bb2fda13600c497272a8dd029313afdb125fd3

func mooniswap_handler(c echo.Context) error {

	// from := c.FormValue("from")
	// to := c.FormValue("to")
	amount := c.FormValue("amount")

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "amount err: amount should be numeric")
	}

	// TODO: Add router
	mooniswapUSDCDaiAddr := common.HexToAddress("0x31631b3dd6c697e574d6b886708cd44f5ccf258f")
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("cannot connect infura"))
	}

	mooniswapModule, err := contractabi.NewMooniswap(mooniswapUSDCDaiAddr, conn)
	if err != nil {
		c.Logger().Error("2")

		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := mooniswapModule.GetReturn(nil, common.HexToAddress(m1["USDC"].Address), common.HexToAddress(m1["DAI"].Address), big.NewInt(int64(s)))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}
