package handler

import (
	contractabi "aggregator_info/contract_abi"
	"aggregator_info/types"
	"errors"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Uniswap_v2_handler(from, to, amount string) (*types.ExchangePair, error) {

	UniswapV2Result := new(types.ExchangePair)
	UniswapV2Result.ContractName = "UniswapV2"

	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return UniswapV2Result, errors.New("amount err: amount should be numeric")
	}

	uniswapV2Addr := common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
	if err != nil {
		return UniswapV2Result, errors.New("cannot connect infura")
	}
	defer conn.Close()

	uniswapV2Module, err := contractabi.NewUniswapV2(uniswapV2Addr, conn)

	if err != nil {
		return UniswapV2Result, err
	}
	path := make([]common.Address, 2)
	if from == "ETH" {
		from = "WETH"
	}
	if to == "ETH" {
		to = "WETH"
	}
	path[0] = common.HexToAddress(M1[from].Address)
	path[1] = common.HexToAddress(M1[to].Address)
	result, err := uniswapV2Module.GetAmountsOut(nil, big.NewInt(int64(s)), path)
	if err != nil {
		return UniswapV2Result, err
	}
	UniswapV2Result.Ratio = result[1].String()

	return UniswapV2Result, nil
}

// func ReadABIFile(filePath string) string {
// 	f, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		fmt.Println("read fail", err)
// 	}
// 	return string(f)
// }

// func estimate_uniswap_gas() (uint64, error) {

// 	// addr of Uniswap
// 	address := common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")

// 	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688")
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer conn.Close()

// 	RawABI := ReadABIFile("raw_abi/uniswapv2.abi")
// 	parsed, err := abi.JSON(strings.NewReader(RawABI))
// 	if err != nil {
// 		return 0, err
// 	}

// 	valueInput, err := parsed.Pack(
// 		"swapExactETHForTokens",
// 		big.NewInt(10000),
// 		[]common.Address{common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"), common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"), common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")},
// 		common.HexToAddress("0xcd69c8CbBFe5b1219C0f8911204aA961294E74e3"),
// 		big.NewInt(time.Now().Unix()+600),
// 	)
// 	if err != nil {
// 		return 0, err
// 	}
// 	fmt.Println("packed result:", fmt.Sprintf("%x", valueInput))

// 	// Get suggestGas
// 	suggestGas, _ := conn.SuggestGasPrice(context.Background())
// 	fmt.Println("SuggestedGasPrice is:", suggestGas)

// 	// estimate gas used when call swapExactETHForTokens

// 	rawData := "7ff36ab500000000000000000000000000000000000000000000000032576437433d3a000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000cd69c8cbbfe5b1219c0f8911204aa961294e74e3000000000000000000000000000000000000000000000000000000005f5c86d00000000000000000000000000000000000000000000000000000000000000003000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000df5e0e81dff6faf3a7e52ba697820c5e32d806a80000000000000000000000006b175474e89094c44da98b954eedeac495271d0f"
// 	bytesData, err := hex.DecodeString(rawData)
// 	if err != nil {
// 		return 0, err
// 	}

// 	estimateGas, err := conn.EstimateGas(context.Background(), ethereum.CallMsg{
// 		From: address,
// 		To:   &address,
// 		// Data: valueInput,
// 		Data: bytesData,
// 	})
// 	if err != nil {
// 		fmt.Println("EstimateGas Err: ", err)
// 		return 0, err
// 	}

// 	return estimateGas, nil
// }
