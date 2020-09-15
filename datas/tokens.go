package datas

import "aggregator_info/types"

// TokenInfos addr of ERC20 tokens
var TokenInfos = make(map[string]types.Token)

func init() {
	ConstructToken()
}

// ConstructToken addr of ERC20 tokens
func ConstructToken() {
	TokenInfos["ETH"] = types.Token{
		Name:    "ETH",
		Address: "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
	}

	TokenInfos["USDC"] = types.Token{
		Name:    "USDC",
		Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	}
	TokenInfos["YFI"] = types.Token{
		Name:    "YFI",
		Address: "0x0bc529c00c6401aef6d220be8c6ea1667f6ad93e",
	}
	TokenInfos["CRV"] = types.Token{
		Name:    "CRV",
		Address: "0xd533a949740bb3306d119cc777fa900ba034cd52",
	}
	TokenInfos["BZRX"] = types.Token{
		Name:    "BZRX",
		Address: "0x56d811088235f11c8920698a204a5010a788f4b3",
	}
	TokenInfos["YFII"] = types.Token{
		Name:    "YFII",
		Address: "0xa1d0e215a23d7030842fc67ce582a6afa3ccab83",
	}
	TokenInfos["SRM"] = types.Token{
		Name:    "SRM",
		Address: "0x476c5e26a75bd202a9683ffd34359c0cc15be0ff",
	}
	TokenInfos["ANJ"] = types.Token{
		Name:    "ANJ",
		Address: "0xcd62b1c403fa761baadfc74c525ce2b51780b184",
	}
	TokenInfos["YAMv2"] = types.Token{
		Name:    "YAMv2",
		Address: "0xaba8cac6866b83ae4eec97dd07ed254282f6ad8a",
	}
	TokenInfos["YFV"] = types.Token{
		Name:    "YFV",
		Address: "0x45f24baeef268bb6d63aee5129015d69702bcdfa",
	}
	TokenInfos["CVP"] = types.Token{
		Name:    "CVP",
		Address: "0x38e4adb44ef08f22f5b5b76a8f0c2d0dcbe7dca1",
	}
	TokenInfos["KIMCHI"] = types.Token{
		Name:    "KIMCHI",
		Address: "0x1e18821e69b9faa8e6e75dffe54e7e25754beda0",
	}
	TokenInfos["YUNO"] = types.Token{
		Name:    "YUNO",
		Address: "0x4b4f5286e0f93e965292b922b9cd1677512f1222",
	}
	TokenInfos["SUSHI"] = types.Token{
		Name:    "SUSHI",
		Address: "0x6b3595068778dd592e39a122f4f5a5cf09c90fe2",
	}
	TokenInfos["FARM"] = types.Token{
		Name:    "FARM",
		Address: "0xa0246c9032bc3a600820415ae600c6388619a14d",
	}
	TokenInfos["USDT"] = types.Token{
		Name:    "USDT",
		Address: "0xdac17f958d2ee523a2206206994597c13d831ec7",
	}
	TokenInfos["LINK"] = types.Token{
		Name:    "LINK",
		Address: "0x514910771af9ca656af840dff83e8264ecf986ca",
	}
	TokenInfos["DAI"] = types.Token{
		Name:    "DAI",
		Address: "0x6b175474e89094c44da98b954eedeac495271d0f",
	}
	TokenInfos["BAT"] = types.Token{
		Name:    "BAT",
		Address: "0x0d8775f648430679a709e98d2b0cb6250d2887ef",
	}
	TokenInfos["KNC"] = types.Token{
		Name:    "KNC",
		Address: "0xdd974d5c2e2928dea5f71b9825b8b646686bd200",
	}
	TokenInfos["WETH"] = types.Token{
		Name:    "WETH",
		Address: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
	}

	TokenInfos["AMPL"] = types.Token{
		Name:    "AMPL",
		Address: "0xd46ba6d942050d489dbd938a2c909a5d5039a161",
	}

	TokenInfos["COMP"] = types.Token{
		Name:    "COMP",
		Address: "0xc00e94cb662c3520282e6f5717214004a7f26888",
	}

	TokenInfos["LEND"] = types.Token{
		Name:    "LEND",
		Address: "0x80fb784b7ed66730e8b1dbd9820afd29931aab03",
	}

	TokenInfos["SNX"] = types.Token{
		Name:    "SNX",
		Address: "0xc011a73ee8576fb46f5e1c5751ca3b9fe0af2a6f",
	}

	TokenInfos["MKR"] = types.Token{
		Name:    "MKR",
		Address: "0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2",
	}

	TokenInfos["BAL"] = types.Token{
		Name:    "BAL",
		Address: "0xba100000625a3754423978a60c9317c58a424e3d",
	}

	TokenInfos["WBTC"] = types.Token{
		Name:    "WBTC",
		Address: "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599",
	}

	TokenInfos["LRC"] = types.Token{
		Name:    "LRC",
		Address: "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd",
	}

	TokenInfos["RPL"] = types.Token{
		Name:    "RPL",
		Address: "0xb4efd85c19999d84251304bda99e90b92300bd93",
	}

	TokenInfos["YFFI"] = types.Token{
		Name:    "YFFI",
		Address: "0xcee1d3c3a02267e37e6b373060f79d5d7b9e1669",
	}
}
