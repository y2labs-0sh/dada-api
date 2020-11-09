// +build !testnet

package data

import "strings"

const (
	EtherScanAPIKey  = "RUHXBDW3HQZ7MQMVN2GSZJAFXQS9RDGEP4"
	CryptocompareAPI = "87cf5b483c21a7dd2846c91307882bd00c011a5ca4dd8bc5c529b65666788ec8"
)

// Contract ABI:
const (
	UniswapV2               = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
	Bancor                  = "0x2F9EC37d6CcFFf1caB21733BdaDEdE11c823cCB0"
	OneInch                 = "0xC586BeF4a0992C495Cf22e1aeEE4E446CECDee0E"
	SushiSwap               = "0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F"
	Kyber                   = "0x818E6FECD516Ecc3849DAf6845e3EC868087B755"
	MooniswapFactory        = "0x71CD6666064C3A1354a3B4dca5fA1E2D3ee7D303"
	Paraswap                = "0x86969d29F5fd327E1009bA66072BE22DB6017cC6"
	Oasis                   = "0x794e6e91555438aFc3ccF1c5076A74F42133d08D"
	Dforce                  = "0x03eF3f37856bD08eb47E2dE7ABc4Ddd2c19B60F2"
	UniswapV2Factory        = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
	BalancerRegistry        = "0x65e67cbc342712DF67494ACEfc06fe951EE93982"
	BalancerExchangeProxyV2 = "0x3E66B66Fd1d0b02fDa6C811Da9E0547970DB2f21"
	SushiSwapFactory        = "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac"
)

var (
	ethereumPort = [...]string{
		"http://47.241.27.12:8545/",
		"https://mainnet.infura.io/v3/bd0da7059b9d48e8bf3c92681bbcb636",
		"https://mainnet.infura.io/v3/e468cafc35eb43f0b6bd2ab4c83fa688",
	}

	AddrNameTag = make(map[string]string)
	TxAction    = make(map[string]string)
)

func init() {
	AddrNameTag[strings.ToLower("0x6B175474E89094C44Da98b954EedeAC495271d0F")] = "Dai Stablecoin"
	AddrNameTag[strings.ToLower("0xdAC17F958D2ee523a2206206994597C13D831ec7")] = "Tether: USDT Stablecoin"
	AddrNameTag[strings.ToLower("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")] = "USD Coin"
	AddrNameTag[strings.ToLower("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")] = "Uniswap V2: Router 2"
	AddrNameTag[strings.ToLower("0x9AAb3f75489902f3a48495025729a0AF77d4b11e")] = "Kyber: Proxy 2"
	AddrNameTag[strings.ToLower("0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F")] = "SushiSwap: Router"
	AddrNameTag[strings.ToLower("0x343E3a490c9251dC0eaA81Da146ba6ABe6C78b2d")] = "Zapper.Fi: Uniswap Zap Out V2"
	AddrNameTag[strings.ToLower("0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc")] = "Uniswap V2: USDC 3"
	AddrNameTag[strings.ToLower("0x80C5e6908368CB9db503bA968d7ec5A565BFb389")] = "Zapper.Fi: Uniswap V2 Zap In"
	AddrNameTag[strings.ToLower("0x3E66B66Fd1d0b02fDa6C811Da9E0547970DB2f21")] = "Balancer: Exchange Proxy 2"

	TxAction[strings.ToLower("0x095ea7b3")] = "Approve"
}

func GetEthereumPort() string {
	return ethereumPort[1]
}
