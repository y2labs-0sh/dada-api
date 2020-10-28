package liquidity_history

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

// Get add liquidity history of uniswap platform
// (including 3 contract: dada-uniswap-zapIn-general)

func TestGetUniswapInvestHistory() {
	result, err := GetUniswapInvestHistory(common.HexToAddress("0x0320304825e3a1156c70fc2607a21c130e21583d"))
	if err != nil {
		fmt.Println(err)
	}

	for _, aResult := range result {
		logrus.WithFields(logrus.Fields{
			"Platform":         aResult.Platform,
			"Ops":              aResult.Ops,
			"TxHash":           aResult.TxHash,
			"Timestamp":        aResult.Timestamp,
			"BlockNumber":      aResult.BlockNumber,
			"ContractAddr":     aResult.ContractAddr,
			"FromTokenAddress": aResult.FromTokenAddress.String(),
			"ToUniPoolToken0":  aResult.ToUniPoolToken0.String(),
			"ToUniPoolToken1":  aResult.ToUniPoolToken1.String(),
			"Amount":           aResult.Amount.String(),
			"MinPoolToken":     aResult.MinPoolToken.String(),
		}).Info("Get a Tx ops")
	}
}

// Dada add liquidity ops
type DadaInvestOps struct {
	Platform         string
	Ops              string
	TxHash           string
	Timestamp        string
	BlockNumber      string
	ContractAddr     string
	FromTokenAddress common.Address
	ToUniPoolToken0  common.Address
	ToUniPoolToken1  common.Address
	Amount           *big.Int
	MinPoolToken     *big.Int
}

// Add liquidity to uniswap
func GetUniswapInvestHistory(userAddr common.Address) ([]*DadaInvestOps, error) {

	userTxHistory, err := GetAccountTxHistory(userAddr.String())
	if err != nil {
		return nil, err
	}

	var ops string
	out := []*DadaInvestOps{}

	for _, aTxDetail := range userTxHistory.Result {

		if len(aTxDetail.Input) < 10 {
			continue
		}
		funcHash := strings.ToLower(aTxDetail.Input[:10])

		// Invest (current contract use Invest func)
		// function Invest(address _toWhomToIssue, address _FromTokenContractAddress, address _ToUnipoolToken0, address _ToUnipoolToken1, uint256 _amount, uint256 _minPoolTokens)
		// ZapIn(address _toWhomToIssue, address _FromTokenContractAddress, address _ToUnipoolToken0, address _ToUnipoolToken1, uint256 _amount, uint256 _minPoolTokens)
		if funcHash == "0x41490f2a" || funcHash == "0x1d572320" {
			// Uniswap Invest Addr: 0x99bD6fE8e3b522b8475f070bbc880d731c40e9D8

			if funcHash == "0x41490f2a" {
				ops = "Invest"
			} else {
				ops = "ZapIn"
			}

			fromTokenContractAddr := common.HexToAddress(aTxDetail.Input[74:138])
			toUniPoolToken0 := common.HexToAddress(aTxDetail.Input[138:202])
			toUniPoolToken1 := common.HexToAddress(aTxDetail.Input[202:266])
			amount, _ := big.NewInt(0).SetString(aTxDetail.Input[266:330], 16)
			minPoolToken, _ := big.NewInt(0).SetString(aTxDetail.Input[330:], 16)

			out = append(out, &DadaInvestOps{
				Platform:         "DadaUniswap",
				Ops:              ops,
				TxHash:           aTxDetail.Hash,
				Timestamp:        aTxDetail.TimeStamp,
				BlockNumber:      aTxDetail.BlockNumber,
				ContractAddr:     aTxDetail.ContractAddress,
				FromTokenAddress: fromTokenContractAddr,
				ToUniPoolToken0:  toUniPoolToken0,
				ToUniPoolToken1:  toUniPoolToken1,
				Amount:           amount,
				MinPoolToken:     minPoolToken,
			})
		}

		// record stake LP ops: stake: Stake(amount) 0xa694fc3a; withdraw(amount): 0x2e1a7d4d
		if funcHash == "0xa694fc3a" || funcHash == "0x2e1a7d4d" {
			if funcHash == "0xa694fc3a" {
				ops = "Stake"
			} else {
				ops = "Withdraw"
			}

			amount, _ := big.NewInt(0).SetString(aTxDetail.Input[10:74], 16)
			out = append(out, &DadaInvestOps{
				Platform:     "DadaUniswap",
				Ops:          ops,
				TxHash:       aTxDetail.Hash,
				Timestamp:    aTxDetail.TimeStamp,
				BlockNumber:  aTxDetail.BlockNumber,
				ContractAddr: aTxDetail.ContractAddress,
				Amount:       amount,
			})
		}

		// record exit staked LP: Exit() -> 0xe9fad8ee
		if funcHash == "0xe9fad8ee" {
			out = append(out, &DadaInvestOps{
				Platform:     "DadaUniswap",
				Ops:          "Exit",
				TxHash:       aTxDetail.Hash,
				Timestamp:    aTxDetail.TimeStamp,
				BlockNumber:  aTxDetail.BlockNumber,
				ContractAddr: aTxDetail.ContractAddress,
				// TODO: add amount
			})
		}
	}
	return out, nil
}
