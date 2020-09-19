package estimatetxrate

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/y2labs-0sh/aggregator_info/contract_abi"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// UniswapV2Handler get token exchange rate based on from amount
func UniswapV2Handler(from, to string, amount *big.Int) (*types.ExchangePair, error) {

	UniswapV2Result := new(types.ExchangePair)

	if from == "ETH" {
		from = "WETH"
	}
	if to == "ETH" {
		to = "WETH"
	}

	uniswapV2Addr := common.HexToAddress(data.UniswapV2)
	conn, err := ethclient.Dial(fmt.Sprintf(data.InfuraAPI, data.InfuraKey))
	if err != nil {
		return UniswapV2Result, err
	}
	defer conn.Close()

	uniswapV2Module, err := contractabi.NewUniswapV2(uniswapV2Addr, conn)
	if err != nil {
		return UniswapV2Result, err
	}

	path := make([]common.Address, 2)
	path[0] = common.HexToAddress(data.TokenInfos[from].Address)
	path[1] = common.HexToAddress(data.TokenInfos[to].Address)

	result, err := uniswapV2Module.GetAmountsOut(nil, amount, path)
	if err != nil {
		return UniswapV2Result, err
	}

	result[1] = result[1].Mul(result[1], big.NewInt(int64(math.Pow10((18 - int(data.TokenInfos[to].Decimals))))))

	UniswapV2Result.ContractName = "UniswapV2"
	UniswapV2Result.Ratio = result[1].String()
	UniswapV2Result.TxFee = estimatetxfee.TxFeeOfContract["UniswapV2"]
	UniswapV2Result.SupportSwap = true

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
