package harvestfarm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	KEY_API_UI = "41e90ced-d559-4433-b390-af424fdc76d6"
	KEY_API    = "fc8ad696-7905-4daa-a552-129ede248e33"
	API_UI     = "https://api-ui.harvest.finance/%s?key=%s"
	API        = "https://api.harvest.finance/%s?key=%s"

	VAULTS = "vaults"
	POOLS  = "pools"
)

type LPToken struct {
	Address  string      `json:"address"`
	Decimals string      `json:"decimals"`
	Symbol   string      `json:"symbol"`
	Price    interface{} `json:"price"`
}

type Pool struct {
	ID                     string   `json:"id"`
	DisplayName            string   `json:"displayName"`
	Logo                   string   `json:"logo"`
	TokenForLogo           string   `json:"tokenForLogo"`
	Type                   int      `json:"type"`
	ContractAddress        string   `json:"contractAddress"`
	AutoStakePoolAddress   string   `json:"autoStakePoolAddress"`
	CollateralAddress      string   `json:"collateralAddress"`
	RewardAPY              string   `json:"rewardAPY"`
	RewardToken            string   `json:"rewardToken"`
	RewardTokenSymbols     []string `json:"rewardTokenSymbols"`
	CollateralTokenSymbols []string `json:"collateralTokenSymbols"`
	DetailsTokenSymbol     string   `json:"detailsTokenSymbol"`
	LpTokenData            LPToken  `json:"lpTokenData"`
	RewardAPR              string   `json:"rewardAPR"`
	RewardPerToken         string   `json:"rewardPerToken"`
	TotalSupply            string   `json:"totalSupply"`
	FinishTime             string   `json:"finishTime"`
	TotalValueLocked       string   `json:"totalValueLocked"`
	StakeDisabled          bool     `json:"stakeDisabled"`

	Vault *Vault `json:"vault"`
}

type Vault struct {
	LogoURL                         string   `json:"logoUrl"`
	ApyIconURL                      string   `json:"apyIconUrl"`
	TokenAddress                    string   `json:"tokenAddress"`
	Decimals                        string   `json:"decimals"`
	VaultAddress                    string   `json:"vaultAddress"`
	CMCRewardTokenSymbols           []string `json:"cmcRewardTokenSymbols"`
	PricePerFullShare               string   `json:"pricePerFullShare"`
	EstimatedAPY                    string   `json:"estimatedApy"`
	UnderlyingBalanceWithInvestment string   `json:"underlyingBalanceWithInvestment"`
	USDPrice                        string   `json:"usdPrice"`
}

type Vaults map[string]Vault
type Pools []Pool

func GetCurrentPools() (Pools, error) {
	resp, err := http.Get(fmt.Sprintf(API_UI, POOLS, KEY_API_UI))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := make(Pools, 0)
	if err := json.Unmarshal(respBytes, &res); err != nil {
		return nil, err
	}
	effectivePools := make(Pools, 0)
	for _, v := range res {
		if !v.StakeDisabled {
			effectivePools = append(effectivePools, v)
		}
	}
	return effectivePools, nil
}

func GetVaults() (Vaults, error) {
	resp, err := http.Get(fmt.Sprintf(API_UI, VAULTS, KEY_API_UI))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	vaults := make(Vaults)
	if err := json.Unmarshal(respBytes, &vaults); err != nil {
		return nil, err
	}
	return vaults, nil
}
