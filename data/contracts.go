// +build !testnet

package data

const (
	EtherScanAPIKey = "RUHXBDW3HQZ7MQMVN2GSZJAFXQS9RDGEP4"
	InfuraAPI       = "https://mainnet.infura.io/v3/%s"
	InfuraKey       = "bd0da7059b9d48e8bf3c92681bbcb636"
	// InfuraKey       = "e468cafc35eb43f0b6bd2ab4c83fa688"

)

// Contract ABI:
const (
	UniswapV2               = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
	Bancor                  = "0x2F9EC37d6CcFFf1caB21733BdaDEdE11c823cCB0"
	OneInch                 = "0xC586BeF4a0992C495Cf22e1aeEE4E446CECDee0E"
	SushiSwap               = "0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F"
	Kyber                   = "0x818E6FECD516Ecc3849DAf6845e3EC868087B755"
	MooniswapFactory        = "0x71CD6666064C3A1354a3B4dca5fA1E2D3ee7D303"
	Paraswap                = "0x86969d29F5fd327E1009bA66072BE22DB6017cC6" // TODO: 来自estimate_tx_fee package
	Oasis                   = "0x794e6e91555438aFc3ccF1c5076A74F42133d08D"
	Dforce                  = "0x03eF3f37856bD08eb47E2dE7ABc4Ddd2c19B60F2"
	Paraswap2               = "0x12295f06DA62693086F5DA45b78e20B778060853" // 来自Handler package
	UniswapV2Factory        = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
	BalancerRegistry        = "0x65e67cbc342712DF67494ACEfc06fe951EE93982"
	BalancerExchangeProxyV2 = "0x3E66B66Fd1d0b02fDa6C811Da9E0547970DB2f21"
)

// UniswapV2
// "7ff36ab5" swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline)
// "791ac947" swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)
// "fb3bdb41" swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline)
// "38ed1739" swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)
// "4a25d94a" swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline)
// "18cbafe5" swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)

// Bancor
// `0xe57738e5` claimAndConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee)
// `0x569706eb` convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee)
// `0xb77d239b` convertByPath(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee)

// 1inch
// `0xe2a7515e` swap(address fromToken, address toToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 featureFlags)

// SushiSwap
// `0x18cbafe5` swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)
// `0x7ff36ab5` swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline)
// `0x38ed1739` swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)

// Kyber
// `0xcb3c28c7` trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId)
// `0x29589f61` tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint)

// Paraswap
// `c5f0b909`: function swap(IERC20 fromToken, IERC20 toToken,uint256 fromAmount,uint256 toAmount,address exchange,bytes calldata payload) external payable returns (uint256);

// Oasis
// `0x1b33d412` offer(uint256 pay_amt, address pay_gem, uint256 buy_amt, address buy_gem, uint256 pos)

// Dforce
// `0xdf791e50` swap(address source, address dest, uint256 sourceAmount)
