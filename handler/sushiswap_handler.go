package handler

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

func Sushiswap_handler(c echo.Context) error {

	amount := c.FormValue("amount")
	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "amount err: amount should be numeric")
	}

	sushiSwapAddr := common.HexToAddress("0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F") // TODO: change addr
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("cannot connect infura"))
	}

	sushiSwapModule, err := contractabi.NewSushiSwap(sushiSwapAddr, conn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	path := make([]common.Address, 2)
	path[0] = common.HexToAddress(M1["DAI"].Address)
	path[1] = common.HexToAddress(M1["USDT"].Address)
	result, err := sushiSwapModule.GetAmountsOut(nil, big.NewInt(int64(s)), path)
	if err != nil {
		c.Logger().Error("2")
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, result)
}
