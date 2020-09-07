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

func Uniswap_v2_handler(c echo.Context) error {
	amount := c.FormValue("amount")

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "amount err: amount should be numeric")
	}

	uniswapV2Addr := common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("cannot connect infura"))
	}

	uniswapV2Module, err := contractabi.NewUniswapV2(uniswapV2Addr, conn)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	path := make([]common.Address, 2)
	path[0] = common.HexToAddress(M1["USDC"].Address)
	path[1] = common.HexToAddress(M1["DAI"].Address)
	result, err := uniswapV2Module.GetAmountsOut(nil, big.NewInt(int64(s)), path)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, result)
}
