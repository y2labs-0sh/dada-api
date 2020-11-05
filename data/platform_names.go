package data

const (
	uniswap_name        = "UniswapV2"
	harvest_invest_name = "HarvestInvest"
	harvest_reward_name = "HarvestReward"
	sushiswap_name      = "Sushiswap"
	balancer_name       = "Balancer"
)

type dexNames struct {
	Uniswap       string
	HarvestInvest string
	HarvestReward string
	Sushiswap     string
	Balancer      string
}

var dexnames = dexNames{
	Uniswap:       uniswap_name,
	HarvestInvest: harvest_invest_name,
	HarvestReward: harvest_reward_name,
	Sushiswap:     sushiswap_name,
	Balancer:      balancer_name,
}

func DexNames() dexNames {
	return dexnames
}
