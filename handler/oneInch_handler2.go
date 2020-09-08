package handler

import (
	contractabi "aggregator_info/contract_abi"
	"aggregator_info/types"
	"errors"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"
)

func OneInch_handler(c echo.Context) error {
	from := c.FormValue("from")
	to := c.FormValue("to")
	amount := c.FormValue("amount")

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "amount err: amount should be numeric")
	}

	oneInchModuleAddr := common.HexToAddress("0xC586BeF4a0992C495Cf22e1aeEE4E446CECDee0E")
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("cannot connect infura"))
	}
	defer conn.Close()

	oneInchModule, err := contractabi.NewOneinch(oneInchModuleAddr, conn)

	result, err := oneInchModule.GetExpectedReturn(nil, common.HexToAddress(M1[from].Address), common.HexToAddress(M1[to].Address), big.NewInt(int64(s)), big.NewInt(10), big.NewInt(0))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	result2 := types.Exchange_pair{
		FromName: from,
		ToName:   to,
		FromAddr: M1[from].Address,
		ToAddr:   M1[to].Address,
		Ratio:    result.ReturnAmount.String(),
	}

	return c.JSON(http.StatusOK, result2)
}
