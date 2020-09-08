package handler

import (
	contractabi "aggregator_info/contract_abi"
	"aggregator_info/types"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"
)

func getFactory(token1, token2 string) (string, error) {
	mooniswapFactoryAddr := common.HexToAddress("0x71CD6666064C3A1354a3B4dca5fA1E2D3ee7D303")
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	mooniswapFactoryModule, err := contractabi.NewMooniswapFactory(mooniswapFactoryAddr, conn)
	if err != nil {
		return "", err
	}

	poolAddr, err := mooniswapFactoryModule.Pools(nil, common.HexToAddress(M1[token1].Address), common.HexToAddress(M1[token2].Address))

	if err != nil {
		return "", err
	}
	return poolAddr.String(), nil
}

func Mooniswap_handler(c echo.Context) error {

	// eg: from: USDC, to: DAI
	from := c.FormValue("from")
	to := c.FormValue("to")
	amount := c.FormValue("amount")

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "amount err: amount should be numeric")
	}

	poolAddr, err := getFactory(from, to)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	mooniswapUSDCDaiAddr := common.HexToAddress(poolAddr)
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("cannot connect infura"))
	}
	defer conn.Close()

	mooniswapModule, err := contractabi.NewMooniswap(mooniswapUSDCDaiAddr, conn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := mooniswapModule.GetReturn(nil, common.HexToAddress(M1[from].Address), common.HexToAddress(M1[to].Address), big.NewInt(int64(s)))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	result2 := types.Exchange_pair{
		FromName: from,
		ToName:   to,
		FromAddr: M1[from].Address,
		ToAddr:   M1[to].Address,
		Ratio:    fmt.Sprintf("%v", result.Int64()),
	}

	return c.JSON(http.StatusOK, result2)
}
