package main

type Token struct {
	Name    string `json:"name"`
	Address string `json:"addr"`
}

type Exchange struct {
	Name       string `json:"name"`
	APIAddress string `json:"api_address"`
	Decimals   int    `json:"decimals"` // TODO: 写上每个代币的Decimals?
}

type ReturnResults struct {
	ExchangePair []exchange_pair `json:"exchange_pair"`
}

type exchange_pair struct {
	FromName string `json:"from_name"`
	ToName   string `json:"to_name"`
	FromAddr string `json:"from_addr"`
	ToAddr   string `json:"to_addr"`
	Ratio    string `json:"ratio"`
	TxGas    string `json:"tx_gas"`
}

func constructExchange() {

	// API doc
	m2["1inch.exchange"] = Exchange{
		Name:       "",
		APIAddress: "https://api.1inch.exchange/v1.1/quote?fromTokenSymbol=ETH&toTokenSymbol=DAI&amount=100000000000000000000",
	}

	// TODO: 查询第一个
	m2["Uniswap V2"] = Exchange{
		Name:       "Uniswap V2",
		APIAddress: "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2",
	}

	m2["Mooniswap"] = Exchange{
		Name:       "Mooniswap",
		APIAddress: "",
	}
	m2["Bancor"] = Exchange{
		Name:       "Bancor",
		APIAddress: "",
	}
	m2["0x"] = Exchange{
		Name:       "0x",
		APIAddress: "",
	}
	m2["ParaSwap"] = Exchange{
		Name:       "ParaSwap",
		APIAddress: "",
	}
	m2["Uniswap"] = Exchange{
		Name:       "Uniswap",
		APIAddress: "",
	}
	m2["dForce Swap"] = Exchange{
		Name:       "dForce Swap",
		APIAddress: "",
	}
	m2["Kyber"] = Exchange{
		Name:       "Kyber",
		APIAddress: "",
	}
	m2["Oasis"] = Exchange{
		Name:       "Oasis",
		APIAddress: "",
	}
	m2["dYdX"] = Exchange{
		Name:       "dYdX",
		APIAddress: "",
	}
	m2["Curve"] = Exchange{
		Name:       "Curve",
		APIAddress: "",
	}
	m2["DDEX"] = Exchange{
		Name:       "DDEX",
		APIAddress: "",
	}
	m2["Balancer"] = Exchange{
		Name:       "Balancer",
		APIAddress: "",
	}
	m2["Loopring"] = Exchange{
		Name:       "Loopring",
		APIAddress: "",
	}

	m2["DODO"] = Exchange{
		Name:       "DODO",
		APIAddress: "",
	}

	m2["DEX.AG"] = Exchange{
		Name:       "DEX.AG",
		APIAddress: "",
	}

	m2[""] = Exchange{
		Name:       "",
		APIAddress: "",
	}

	m2["Tokenlon"] = Exchange{
		Name:       "Tokenlon",
		APIAddress: "",
	}
	m2["IDEX"] = Exchange{
		Name:       "IDEX",
		APIAddress: "",
	}
}

func constructToken() {
	m1["ETH"] = Token{
		Name:    "ETH",
		Address: "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
	}

	m1["USDC"] = Token{
		Name:    "USDC",
		Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	}
	m1["YFI"] = Token{
		Name:    "YFI",
		Address: "0x0bc529c00c6401aef6d220be8c6ea1667f6ad93e",
	}
	m1["CRV"] = Token{
		Name:    "CRV",
		Address: "0xd533a949740bb3306d119cc777fa900ba034cd52",
	}
	m1["BZRX"] = Token{
		Name:    "BZRX",
		Address: "0x56d811088235f11c8920698a204a5010a788f4b3",
	}
	m1["YFII"] = Token{
		Name:    "YFII",
		Address: "0xa1d0e215a23d7030842fc67ce582a6afa3ccab83",
	}
	m1["SRM"] = Token{
		Name:    "SRM",
		Address: "0x476c5e26a75bd202a9683ffd34359c0cc15be0ff",
	}
	m1["ANJ"] = Token{
		Name:    "ANJ",
		Address: "0xcd62b1c403fa761baadfc74c525ce2b51780b184",
	}
	m1["YAMv2"] = Token{
		Name:    "YAMv2",
		Address: "0xaba8cac6866b83ae4eec97dd07ed254282f6ad8a",
	}
	m1["YFV"] = Token{
		Name:    "YFV",
		Address: "0x45f24baeef268bb6d63aee5129015d69702bcdfa",
	}
	m1["CVP"] = Token{
		Name:    "CVP",
		Address: "0x38e4adb44ef08f22f5b5b76a8f0c2d0dcbe7dca1",
	}
	m1["KIMCHI"] = Token{
		Name:    "KIMCHI",
		Address: "0x1e18821e69b9faa8e6e75dffe54e7e25754beda0",
	}
	m1["YUNO"] = Token{
		Name:    "YUNO",
		Address: "0x4b4f5286e0f93e965292b922b9cd1677512f1222",
	}
	m1["SUSHI"] = Token{
		Name:    "SUSHI",
		Address: "0x6b3595068778dd592e39a122f4f5a5cf09c90fe2",
	}
	m1["FARM"] = Token{
		Name:    "",
		Address: "0xa0246c9032bc3a600820415ae600c6388619a14d",
	}
	m1["USDT"] = Token{
		Name:    "",
		Address: "0xdac17f958d2ee523a2206206994597c13d831ec7",
	}
	m1["LINK"] = Token{
		Name:    "",
		Address: "0x514910771af9ca656af840dff83e8264ecf986ca",
	}
	m1["DAI"] = Token{
		Name:    "",
		Address: "0x6b175474e89094c44da98b954eedeac495271d0f",
	}
	m1["BAT"] = Token{
		Name:    "",
		Address: "0x0d8775f648430679a709e98d2b0cb6250d2887ef",
	}
	m1["KNC"] = Token{
		Name:    "",
		Address: "0xdd974d5c2e2928dea5f71b9825b8b646686bd200",
	}
	m1["WETH"] = Token{
		Name:    "",
		Address: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
	}

	m1["WETH"] = Token{
		Name:    "",
		Address: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
	}

	m1["AMPL"] = Token{
		Name:    "",
		Address: "0xd46ba6d942050d489dbd938a2c909a5d5039a161",
	}

	m1["COMP"] = Token{
		Name:    "",
		Address: "0xc00e94cb662c3520282e6f5717214004a7f26888",
	}

	m1["LEND"] = Token{
		Name:    "",
		Address: "0x80fb784b7ed66730e8b1dbd9820afd29931aab03",
	}

	m1["SNX"] = Token{
		Name:    "",
		Address: "0xc011a73ee8576fb46f5e1c5751ca3b9fe0af2a6f",
	}

	m1["MKR"] = Token{
		Name:    "",
		Address: "0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2",
	}

	m1["BAL"] = Token{
		Name:    "",
		Address: "0xba100000625a3754423978a60c9317c58a424e3d",
	}

	m1["WBTC"] = Token{
		Name:    "",
		Address: "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599",
	}

	m1["LRC"] = Token{
		Name:    "",
		Address: "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd",
	}

	m1["RPL"] = Token{
		Name:    "",
		Address: "0xb4efd85c19999d84251304bda99e90b92300bd93",
	}

	m1["YFFI"] = Token{
		Name:    "",
		Address: "0xcee1d3c3a02267e37e6b373060f79d5d7b9e1669",
	}
}
