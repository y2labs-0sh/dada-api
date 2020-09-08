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

// https://github.com/curvefi/curve-vue/blob/master/src/docs/README.md
func Curve_handler(c echo.Context) error {

	from := c.FormValue("from")
	to := c.FormValue("to")
	amount := c.FormValue("amount")

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "amount err: amount should be numeric")
	}

	curveAddr, curveToken1, curveToken2, err := curve_router(from, to)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("cannot connect infura"))
	}
	defer conn.Close()

	curveModule, err := contractabi.NewCurve(common.HexToAddress(curveAddr), conn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := curveModule.GetDyUnderlying(nil, big.NewInt(curveToken1), big.NewInt(curveToken2), big.NewInt(int64(s)))

	fromCoinAddr, err := curveModule.Coins(nil, big.NewInt(curveToken1))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	toCoinAddr, err := curveModule.Coins(nil, big.NewInt(curveToken2))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result2 := types.Exchange_pair{
		FromName: from,
		ToName:   to,
		FromAddr: fromCoinAddr.String(),
		ToAddr:   toCoinAddr.String(),
		Ratio:    fmt.Sprintf("%v", result.Int64()),
	}

	return c.JSON(http.StatusOK, result2)
}

// TODO: record those coin's addr
func curve_router(from, to string) (string, int64, int64, error) {
	a1 := "0xA2B47E3D5c44877cca798226B7B8118F9BFb7A56" // DAI, USDC
	a2 := "0x52EA46506B9CC5Ef470C5bf89f17Dc28bB35D85C" // DAI, USDC, USDT
	a3 := "0x06364f10B501e868329afBc005b3492902d6C763" // DAI, USDC, USDT, PAX

	// y-pool
	a4 := "0x45F783CCE6B7FF23B2ab2D70e416cdb7D6055f51" // DAI, USDC, USDT, TUSD
	a5 := "0x79a8C46DeA5aDa233ABaFFD40F3A0A2B1e5A4F27" // DAI, USDC, USDT, BUSD
	a6 := "0xA5407eAE9Ba41422680e2e00537571bcC53efBfD" // DAI, USDC, USDT, sUSD
	a7 := "0x93054188d876f558f4a66B2EF1d97d16eDf0895B" // renBTC, wBTC
	a8 := "0x7fC77b5c7614E1533320Ea6DDc2Eb61fa00A9714" // renBTC, wBTC, sBTC
	fmt.Println(a1, a2, a3, a4, a5, a6, a7, a8)
	return a1, 0, 1, nil
}
