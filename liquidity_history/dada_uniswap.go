package liquidity_history

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/y2labs-0sh/dada-api/invest_factory"
)

// Get add liquidity history of uniswap platform
// (including 3 contract: dada-uniswap-zapIn-general)

func TestGetUniswapInvestHistory() error {
	GetUniswapInvestHistory(common.HexToAddress("0x320304825e3a1156c70fc2607a21c130e21583d"))
	return nil
}

// Dada add liquidity ops
type DadaInvestOps struct {
	Platform         string
	Ops              string
	FromTokenAddress common.Address
	ToUniPoolToken0  common.Address
	ToUniPoolToken1  common.Address
	Amount           *big.Int
	MinPoolToken     *big.Int
}

func GetUniswapInvestHistory(userAddr common.Address) ([]*DadaInvestOps, error) {

	userTxHistory, err := GetAccountTxHistory(userAddr.String())
	if err != nil {
		return nil, err
	}

	out := []*DadaInvestOps{}

	for _, aTxDetail := range userTxHistory.Result {
		// Uniswap Invest Addr: 0x99bD6fE8e3b522b8475f070bbc880d731c40e9D8
		if strings.ToLower(aTxDetail.To) == strings.ToLower(invest_factory.UniswapInvestAddress.String()) {
			if len(aTxDetail.Input) == 394 && aTxDetail.IsError == "0" {

				funcHash := aTxDetail.Input[:10]

				// Invest (current contract use Invest func)
				// function Invest(address _toWhomToIssue, address _FromTokenContractAddress, address _ToUnipoolToken0, address _ToUnipoolToken1, uint256 _amount, uint256 _minPoolTokens)
				// ZapIn(address _toWhomToIssue, address _FromTokenContractAddress, address _ToUnipoolToken0, address _ToUnipoolToken1, uint256 _amount, uint256 _minPoolTokens)
				if funcHash == "0x41490f2a" || funcHash == "0x1d572320" {
					// toWhomToIssue := common.HexToAddress(aTxDetail.Input[10:74])
					fromTokenContractAddr := common.HexToAddress(aTxDetail.Input[74:138])
					toUniPoolToken0 := common.HexToAddress(aTxDetail.Input[138:202])
					toUniPoolToken1 := common.HexToAddress(aTxDetail.Input[202:266])
					amount, _ := big.NewInt(0).SetString(aTxDetail.Input[266:330], 16)
					minPoolToken, _ := big.NewInt(0).SetString(aTxDetail.Input[330:], 16)

					out = append(out, &DadaInvestOps{
						Platform:         "DadaUniswap",
						FromTokenAddress: fromTokenContractAddr,
						ToUniPoolToken0:  toUniPoolToken0,
						ToUniPoolToken1:  toUniPoolToken1,
						Amount:           amount,
						MinPoolToken:     minPoolToken,
					})
				}

			}

			// invest_factory.UniswapInvestAddress

		}
	}
	return out, nil
}
