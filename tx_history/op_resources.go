package tx_history

import "strings"

var (
	ContractName = make(map[string]string)
	OpsName      = make(map[string]string)
)

func init() {

	// Uniswap
	ContractName["0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"] = "Uniswap" // "Uniswap V2: Router 2"
	// Uniswap-Swap:
	OpsName["0x5c11d795"] = "SwapTokenForToken" // swapExactTokensForTokensSupportingFeeOnTransferTokens
	OpsName["0x18cbafe5"] = "SwapTokensForETH"  // swapExactTokensForETH
	OpsName["0x38ed1739"] = "SwapTokenForToken" // swapExactTokensForTokens
	OpsName["0x7ff36ab5"] = "SwapETHForToken"   // swapExactETHForTokens
	OpsName["0xe8e33700"] = "AddLiquidity"      // AddLiquidity
	OpsName["0xf305d719"] = "AddLiquidityETH"   // AddLiquidityETH
	OpsName["0xbaa2abde"] = "RemoveLiquidity"
	OpsName["0x02751cec"] = "RemoveLiquidityETH"
	OpsName["0xaf2979eb"] = "RemoveLiquidityETH" // RemoveLiquidityETHSupportingFeeOnTransferTokens
	OpsName["0xded9382a"] = "RemoveLiquidityETH" // RemoveLiquidityETHWithPermit
	OpsName["0x5b0d5984"] = "RemoveLiquidityETH" // RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens
	OpsName["0x2195995c"] = "RemoveLiquidity"    // RemoveLiquidityWithPermit
	OpsName["0xfb3bdb41"] = "SwapETHForTokens"   // SwapETHForExactTokens

	// Uniswap-Stake
	ContractName["0xCA35e32e7926b96A9988f61d510E038108d8068e"] = "Uniswap" // "Uniswap V2: ETH/WBTC UNI Pool"
	ContractName["0xa1484C3aa22a66C62b77E0AE78E15258bd0cB711"] = "Uniswap" // "Uniswap V2: ETH/DAI UNI Pool"
	ContractName["0x7FBa4B8Dc5E7616e59622806932DBea72537A56b"] = "Uniswap" // "Uniswap V2: ETH/USDC UNI Pool"
	ContractName["0x6C3e4cb2E96B01F4b866965A91ed4437839A121a"] = "Uniswap" // "Uniswap V2: ETH/USDT UNI Pool"
	OpsName["0xecd9ba82"] = "stakeWithPermit"
	OpsName["0xe9fad8ee"] = "exit"
	OpsName["0x3d18b912"] = "getReward"
	OpsName["0xa694fc3a"] = "stake"
	OpsName["0x2e1a7d4d"] = "withdraw"

	// Sushiswap
	ContractName["0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F"] = "Sushiswap" // "SushiSwap: Router"

	ContractName["0xc2EdaD668740f1aA35E4D8f227fB8E17dcA888Cd"] = "SushiSwap_Staking" // SushiSwap MasterChef LP Staking Pool"
	OpsName["0xe2bbb158"] = "deposit"
	OpsName["0x441a3e70"] = "withdraw"

	ContractName["0x818E6FECD516Ecc3849DAf6845e3EC868087B755"] = "Kyber" // Kyberswap
	OpsName["0x3bba21dc"] = "swapTokenToEther"
	OpsName["0xcb3c28c7"] = "swap" // trade
	OpsName["0x7a2a0456"] = "swapEtherToToken"

	ContractName["0x3E66B66Fd1d0b02fDa6C811Da9E0547970DB2f21"] = "Balancer" // "BalancerExchangeProxyV2"
	OpsName["0xe2b39746"] = "Swap"                                          // "MultihopBatchSwapExactIn"

	// MooniswapFactory = "0x71CD6666064C3A1354a3B4dca5fA1E2D3ee7D303"
	// Mooniswap will swap token-pair in different init
	OpsName["0xd5bcb9b5"] = "Swap"

	ContractName["0x99bd6fe8e3b522b8475f070bbc880d731c40e9d8"] = "Dada"

	// DadaContract
	// ContractName["0x9113947fa46edf94bb689173a82007f4c355c2c0"] // Unknow
	OpsName["0x41490f2a"] = "Invest"
	OpsName["0x1d572320"] = "ZapIn"

	ContractName["0x80C5e6908368CB9db503bA968d7ec5A565BFb389"] = "Zapper" // Zapper.Fi: Uniswap V2 Zap In

	for k, v := range ContractName {
		ContractName[strings.ToLower(k)] = v
	}
	for k, v := range OpsName {
		OpsName[strings.ToLower(k)] = v
	}
}
