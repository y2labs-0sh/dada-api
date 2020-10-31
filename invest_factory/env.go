// +build !testnet

package invest_factory

import (
	"github.com/ethereum/go-ethereum/common"
)

var (
	// contract with only the Invest method
	UniswapInvestAddress = common.HexToAddress("0x73a6DbA24743Ce32f645FeeD8c95F9e0d7494eb9")
	// contract with only the MultiAssetInvest method
	UniswapMultiInvestAddress = common.HexToAddress("")

	BalancerInvestAddress      = common.HexToAddress("")
	BalancerMultiInvestAddress = common.HexToAddress("")
)
