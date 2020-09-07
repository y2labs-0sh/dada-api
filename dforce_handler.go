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

func dforce_handler(c echo.Context) error {

	amount := c.FormValue("amount")

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "amount err: amount should be numeric")
	}

	dforceAddr := common.HexToAddress("0x03eF3f37856bD08eb47E2dE7ABc4Ddd2c19B60F2")
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("cannot connect infura"))
	}

	dforceModule, err := contractabi.NewDforce(dforceAddr, conn)
	if err != nil {
		c.Logger().Error("2")

		return c.JSON(http.StatusBadRequest, err)
	}
	result, err := dforceModule.GetAmountByInput(nil, common.HexToAddress(m1["USDC"].Address), common.HexToAddress(m1["DAI"].Address), big.NewInt(int64(s)))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, result)
}
