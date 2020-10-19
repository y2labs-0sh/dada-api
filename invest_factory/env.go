// +build !testnet

package invest_factory

import (
	"github.com/ethereum/go-ethereum/common"
)

var (
	// contract with only the Invest method
	UniswapInvestAddress = common.HexToAddress("0x99bD6fE8e3b522b8475f070bbc880d731c40e9D8")
	// contract with only the MultiAssetInvest method
	UniswapMultiInvestAddress = common.HexToAddress("")

	BalancerInvestAddress = common.HexToAddress("")
)
