## Go Generate

```shell
make generate
```

## Security Check

```shell
make security
```

## Build

```shell
make build-testnet
```

```shell
make build
```

## Build from Docker

```shell
docker build -t dada-api .
docker run -it --name dada-api-container dada-api /bin/bash
docker cp dada-api-container:/home/dada-api/dada-api ./
```

## Build from Podman

_make sure you have access to golang.org_

```shell
podman build -t build ./
podman run --env [NO_PROXY=true] --volume $(pwd):/home/dada-api:z build
```

```fish
podman build -t build ./
podman run --env [NO_PROXY=true] --volume (pwd):/home/dada-api:z build
```

## Example

#### POST `/account/balances`

Request payload:

```json
{
  "address": "0x6b175474e89094c44da98b954eedeac495271d0f",
  "tokens": ["WETH", "WBTC", "UNI", "FARM", "BAL"]
}
```

Response:

```json
{
  "WETH": ["1.0", "1000000000000000000"],
  "FARM": ["1.0", "1000000000000000000"],
  "WBTC": ["1.0", "100000000"]
}
```

#### POST `/aggrInfo`

- from: `USDT`
- to: `DAI`
- amount: `1200000`

```json
{
  "from_name": "ETH",
  "to_name": "DAI",
  "from_addr": "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
  "to_addr": "0x6b175474e89094c44da98b954eedeac495271d0f",
  "exchange_pairs": [
    {
      "contract_name": "1inch",
      "ratio": "743934090251009241497",
      "tx_fee": "614915073000000000"
    },
    {
      "contract_name": "Sushiswap",
      "ratio": "743356707489010273018",
      "tx_fee": "84454919861122738"
    },
    {
      "contract_name": "Kyber",
      "ratio": "743356707489010273018",
      "tx_fee": ""
    },
    {
      "contract_name": "Mooniswap",
      "ratio": "742954730453393981837",
      "tx_fee": "153637461231001546"
    },
    {
      "contract_name": "UniswapV2",
      "ratio": "741948272651386670794",
      "tx_fee": "73077349968450170"
    },
    {
      "contract_name": "Bancor",
      "ratio": "284990699382090366116",
      "tx_fee": "34590062323383505"
    },
    {
      "contract_name": "Paraswap",
      "ratio": "0",
      "tx_fee": "88697086144942528"
    },
    {
      "contract_name": "Dforce",
      "ratio": "0",
      "tx_fee": "20592359068442936"
    },
    {
      "contract_name": "ZeroX",
      "ratio": "-9223372036854775808",
      "tx_fee": ""
    }
  ]
}
```

#### Get: `/tokenlist`

```json
[
  {
    "name": "ETH",
    "addr": "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
  },
  {
    "name": "USDC",
    "addr": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
  },
  {
    "name": "YFI",
    "addr": "0x0bc529c00c6401aef6d220be8c6ea1667f6ad93e"
  }
  ...
]
```

#### POST: `/swapInfo`

- `contract`: `UniswapV2`
- `from`: `ETH`
- `to`: `DAI`
- `amount`: `10000000`
- `user`: `0xcd69c8CbBFe5b1219C0f8911204aA961294E74e3` (用户地址)
- `slippage`: `5` (滑点为 0.05%)

```json
{
  "data": "0x7ff36ab500000000000000000000000000000000000000000000000000000000009882f80000000000000000000000000000000000000000000000000000000000000080000000000000000000000000cd69c8cbbfe5b1219c0f8911204aa961294e74e3000000000000000000000000000000000000000000000000000000005f6304ff0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000006b175474e89094c44da98b954eedeac495271d0f",
  "tx_fee": "76880097964862601",
  "contract_addr": "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",
  "from_token_amount": "10000000",
  "to_token_amount": "3750340403",
  "from_token_addr": "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
  "allowance": "0",
  "allowance_satisfied": true,
  "allowance_data": "0x095ea7b300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000989680"
}
```

- 当 `from` 为 ETH 时，不用去执行 approve

- 当`from`为其他 ERC20 token 时，如果 allowance_satisfied 为 false，需要执行 approve，执行的合约地址为`from_token_addr`, call_data 为：`allowance_data`
- `contract` 可以选择为：`UniswapV2` ，`Bancor`，`Dforce`，`Kyber`，`Mooniswap`，`Sushiswap`

### /invest/\*

#### GET: `/invest/list`

```json
{
  "address": "0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852",
  "platform": "UniswapV2",
  "liquidity": "396157271.3394341760047949599262313",
  "reserves": ["588180.540502564651528039", "198623914.846887"],
  "tokenprices": [
    "0.002961277552886694139943719263862072",
    "337.6920880061331057763689995205936"
  ],
  "volumes": ["6201321.981781429014459259", "2355360952.854604"],
  "reserveUSD": "396157271.3394341760047949599262313",
  "reserveETH": "1176361.081005129303056078",
  "totalSupply": "8.739720560464409088",
  "volumeUSD": "2339564927.901699586387007576320898",
  "tokens": [
    {
      "address": "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
      "name": "Wrapped Ether",
      "symbol": "WETH",
      "logo": "https://1inch.exchange/assets/tokens/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2.png"
    },
    {
      "address": "0xdac17f958d2ee523a2206206994597c13d831ec7",
      "name": "Tether USD",
      "symbol": "USDT",
      "logo": "https://1inch.exchange/assets/tokens/0xdac17f958d2ee523a2206206994597c13d831ec7.png"
    }
  ]
}
```

#### POST: `/invest/estimate`

- dex: "UniswapV2"
- pool: "0x39444e8Ee494c6212054CFaDF67abDBE97e70207"
- amount: "1.6"
- token: "ETH"
- slippage: "300"(滑点为 3%)

```json
{
  "LP": "60881492234176327580",
  "tokens": {
    "MOO": ["1937805605546836037", "1.937805606"],
    "WETH": ["800000000000000000", "0.8"]
  },
  "invest_amount": "1600000000000000000"
}
```

#### POST: `/invest/prepare`

- dex: "UniswapV2"
- pool: "0x39444e8Ee494c6212054CFaDF67abDBE97e70207"
- amount: "1.6"
- token: "ETH"
- slippage: "300"(滑点为 3%)
- user: "0x92E73408801e713f8371f8A8c31a40130ae61a40"

```json
{
  "data": "0x1d57232000000000000000000000000092e73408801e713f8371f8a8c31a40130ae61a400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b93152b59e65a6de8d3464061bcc1d68f6749f98000000000000000000000000c778417e063141139fce010982780140aa0cd5ab00000000000000000000000000000000000000000000000016345785d8a000000000000000000000000000000000000000000000000000000000000000000000",
  "tx_fee": "",
  "contract_addr": "0xf092A8ff521B39af211DC426e8Eb61c0726147A2",
  "from_token_addr": "",
  "from_token_amount": "1600000000000000000",
  "allowance": "",
  "allowance_satisfied": true,
  "allowance_data": ""
}
```

#### POST: `/invest/estimate_prepare`

- dex: "UniswapV2"
- pool: "0x39444e8Ee494c6212054CFaDF67abDBE97e70207"
- amount: "1.6"
- token: "ETH"
- slippage: "300"(滑点为 3%)
- user: "0x92E73408801e713f8371f8A8c31a40130ae61a40"

```json
{
  "prepare": {
    "data": "0x1d57232000000000000000000000000092e73408801e713f8371f8a8c31a40130ae61a400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b93152b59e65a6de8d3464061bcc1d68f6749f98000000000000000000000000c778417e063141139fce010982780140aa0cd5ab00000000000000000000000000000000000000000000000016345785d8a000000000000000000000000000000000000000000000000000000000000000000000",
    "tx_fee": "",
    "contract_addr": "0xf092A8ff521B39af211DC426e8Eb61c0726147A2",
    "from_token_addr": "",
    "from_token_amount": "1600000000000000000",
    "allowance": "",
    "allowance_satisfied": true,
    "allowance_data": ""
  },
  "estimate": {
    "LP": "60881492234176327580",
    "tokens": {
      "MOO": ["1937805605546836037", "1.937805606"],
      "WETH": ["800000000000000000", "0.8"]
    },
    "invest_amount": "1600000000000000000"
  }
}
```

#### POST: `/invest/multiin`

- input without ETH:
  input:
  ```json
  {
    "assets": {"FARM": "10", "USDC": "10"},
    "dex": "UniswapV2",
    "pool": "0x514906fc121c7878424a5c928cad1852cc545892",
    "infinite_allowance": true,
    "user": "0x92e73408801e713f8371f8a8c31a40130ae61a40"
  }
  ```
  output:
  ```json
  {
    "approves": {
      "FARM": {
        "calldata": "0x095ea7b3000000000000000000000000514906fc121c7878424a5c928cad1852cc54589200000000000000ff00000000000000ff00000000000000ff00000000000000ff"
      },
      "USDC": {
        "calldata": "0x095ea7b3000000000000000000000000514906fc121c7878424a5c928cad1852cc54589200000000000000ff00000000000000ff00000000000000ff00000000000000ff"
      }
    },
    "contract_address": "0x99bD6fE8e3b522b8475f070bbc880d731c40e9D8",
    "calldata": "0x12dd066400000000000000000000000092e73408801e713f8371f8a8c31a40130ae61a40000000000000000000000000a0246c9032bc3a600820415ae600c6388619a14d000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000008ac7230489e8000000000000000000000000000000000000000000000000000000000000009896800000000000000000000000000000000000000000000000000000000000000001",
    "tokens": [
      {
        "symbol": "FARM",
        "amount": "43418111851435391",
        "amount_pretty": "0.04341811"
      },
      {"symbol": "USDC", "amount": "10000000", "amount_pretty": "10"}
    ]
  }
  ```
- input with ETH:
  input:
  ```json
  {
    "assets": {"WBTC": "10", "ETH": "10"},
    "dex": "UniswapV2",
    "pool": "0xbb2b8038a1640196fbe3e38816f3e67cba72d940",
    "infinite_allowance": true,
    "user": "0x92e73408801e713f8371f8a8c31a40130ae61a40"
  }
  ```
  output:
  ```json
  {
    "approves": {
      "WBTC": {
        "calldata": "0x095ea7b3000000000000000000000000bb2b8038a1640196fbe3e38816f3e67cba72d94000000000000000ff00000000000000ff00000000000000ff00000000000000ff"
      }
    },
    "contract_address": "0x99bD6fE8e3b522b8475f070bbc880d731c40e9D8",
    "calldata": "0x12dd066400000000000000000000000092e73408801e713f8371f8a8c31a40130ae61a400000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c5990000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003b9aca000000000000000000000000000000000000000000000000008ac7230489e800000000000000000000000000000000000000000000000000000000000000000001",
    "tokens": [
      {"symbol": "WBTC", "amount": "32748681", "amount_pretty": "0.32748681"},
      {"symbol": "ETH", "amount": "10000000000000000000", "amount_pretty": "10"}
    ]
  }
  ```

#### POST: `/invest/remove_liquidity`

- input

```json
{
  "platform": "UniswapV2",
  "lp_address": "0x0d4a11d5EEaaC28EC3F61d100daF4d40471f1852",
  "amount": "7.8099136809e-8",
  "user": "0x92e73408801e713f8371f8a8c31a40130ae61a40"
}
```

- output

```json
{
  "calldata": "0xbaa2abde000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000000000000000000000000000000000122f1241290000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000092e73408801e713f8371f8a8c31a40130ae61a40000000000000000000000000000000000000000000000000000000005fa396f2",
  "contract_address": "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",
  "tokens": null,
  "allowance": "78099136809",
  "allowance_satisfied": false,
  "allowance_data": "0x095ea7b30000000000000000000000007a250d5630b4cf539739df2c5dacb4c659f2488d000000000000000000000000000000000000000000000000000000122f124129"
}
```

### /staking/\*

#### POST: `/staking/stake`

- dex: "UniswapV2"
- pool: "0xa1484C3aa22a66C62b77E0AE78E15258bd0cB711"
- amount: "1.6"
- user: "0x92E73408801e713f8371f8A8c31a40130ae61a40"

```json
{
  "data": "0xa694fc3a0000000000000000000000000000000000000000000000056bc75e2d63100000",
  "tx_fee": "",
  "contract_addr": "0xa1484C3aa22a66C62b77E0AE78E15258bd0cB711",
  "from_token_addr": "0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11",
  "from_token_amount": "100000000000000000000",
  "allowance": "100000000000000000000",
  "allowance_satisfied": false,
  "allowance_data": "0x095ea7b3000000000000000000000000a1484c3aa22a66c62b77e0ae78e15258bd0cb7110000000000000000000000000000000000000000000000056bc75e2d63100000"
}
```

#### GET: `/staking/pools`

```json
[
  "0xCA35e32e7926b96A9988f61d510E038108d8068e",
  "0xa1484C3aa22a66C62b77E0AE78E15258bd0cB711",
  "0x7FBa4B8Dc5E7616e59622806932DBea72537A56b",
  "0x6C3e4cb2E96B01F4b866965A91ed4437839A121a"
]
```

#### POST: `/staking/claim_reward`

- dex: "UniswapV2"
- pool: "0xa1484C3aa22a66C62b77E0AE78E15258bd0cB711"
- user: "0x92E73408801e713f8371f8A8c31a40130ae61a40"

```json
{
  "data": "0xa694fc3a0000000000000000000000000000000000000000000000056bc75e2d63100000",
  "tx_fee": "",
  "contract_addr": "0xa1484C3aa22a66C62b77E0AE78E15258bd0cB711",
  "reward_token_addr": "0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11",
  "reward_amount": "100000000000000000000",
  "reward_amount_pretty": "1.8"
}
```

#### POST: `/staking/withdraw`

- dex: "UniswapV2"
- pool: "0xa1484C3aa22a66C62b77E0AE78E15258bd0cB711"
- amount: "1.6"
- user: "0x92E73408801e713f8371f8A8c31a40130ae61a40"

```json
{
  "data": "0xa694fc3a0000000000000000000000000000000000000000000000056bc75e2d63100000",
  "tx_fee": "",
  "contract_addr": "0xa1484C3aa22a66C62b77E0AE78E15258bd0cB711",
  "staking_token_addr": "0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11",
  "withdraw_amount": "100000000000000000000",
  "withdraw_amount_pretty": "1.6"
}
```

#### POST: `/staking/exit`

- dex: "UniswapV2"
- pool: "0xa1484C3aa22a66C62b77E0AE78E15258bd0cB711"
- user: "0x92E73408801e713f8371f8A8c31a40130ae61a40"

```json
{
  "data": "0xa694fc3a0000000000000000000000000000000000000000000000056bc75e2d63100000",
  "tx_fee": "",
  "contract_addr": "0xa1484C3aa22a66C62b77E0AE78E15258bd0cB711",
  "staking_token_addr": "0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11",
  "reward_token_addr": "0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11",
  "withdraw_amount": "100000000000000000000",
  "reward_amount": "100000000000000000000",
  "withdraw_amount_pretty": "1.8",
  "reward_amount_pretty": "1.8"
}
```

### Harvest Farm

#### GET `/staking/harvestfarm/info`

获取 harvest farm 相关的 vaults 和 pools 信息。会返回全部信息，都是取自 Harvest 官方的 API。选出一个当前 APY 最高的池子作为 Best Pool。

无需参数

`collateralAddress`是池子通过 deposit 获取 fxxxx token 的合约地址，

`contractAddress`是流动性挖矿 FARM 的地址。

APY = rewardAPY + vault.EstimatedApy

```json
{
  "Best": {
    "id": "weth-usdt-uni-farm",
    "displayName": "Uni ETH-USDT",
    "logo": "",
    "tokenForLogo": "ETH-USDT",
    "type": 0,
    "contractAddress": "0x75071F2653fBC902EBaff908d4c68712a5d1C960",
    "autoStakePoolAddress": "",
    "collateralAddress": "0x7ddc3fff0612e75ea5ddc0d6bd4e268f70362cff",
    "rewardAPY": "22.04",
    "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
    "rewardTokenSymbols": ["FARM"],
    "collateralTokenSymbols": ["fUNI_LP_WETH_USDT"],
    "detailsTokenSymbol": "fUNISWAP_LP",
    "lpTokenData": {
      "address": "0x7DDc3ffF0612E75Ea5ddC0d6Bd4e268f70362Cff",
      "decimals": "18",
      "symbol": "fUNI-V2",
      "price": "49915731.60413436545238271767"
    },
    "rewardAPR": "19.958275516794586562",
    "rewardPerToken": "2627789175107516081421",
    "totalSupply": "735338280508354270",
    "finishTime": "1603828471",
    "totalValueLocked": "36772108.0884701346397125958",
    "stakeDisabled": false,
    "vault": {
      "logoUrl": "./icons/eth_usdt.svg",
      "apyIconUrl": "./icons/uni.png",
      "tokenAddress": "0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852",
      "decimals": "18",
      "vaultAddress": "0x7ddc3fff0612e75ea5ddc0d6bd4e268f70362cff",
      "cmcRewardTokenSymbols": ["FARM", "ETH", "USDT"],
      "pricePerFullShare": "1003877225162379827",
      "estimatedApy": "13.78",
      "underlyingBalanceWithInvestment": "738635266929820332",
      "usdPrice": "49813923.70308837487437353461"
    }
  },
  "Type0": [
    {
      "id": "uni-farm-usdc",
      "displayName": "Uniswap Farm",
      "logo": "🦄",
      "tokenForLogo": "",
      "type": 0,
      "contractAddress": "0x99b0d6641A63Ce173E6EB063b3d3AED9A35Cf9bf",
      "autoStakePoolAddress": "",
      "collateralAddress": "0x514906FC121c7878424a5C928cad1852CC545892",
      "rewardAPY": "516.27",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["UNISWAP_LP"],
      "detailsTokenSymbol": "",
      "lpTokenData": {
        "address": "0x514906FC121c7878424a5C928cad1852CC545892",
        "decimals": "18",
        "symbol": "UNI-V2",
        "price": "23041578.71946014578233675405"
      },
      "rewardAPR": "185.068080801035970981",
      "rewardPerToken": "291867125863770829933967",
      "totalSupply": "436309299844668791",
      "finishTime": "1603828191",
      "totalValueLocked": "10070350.07489744436547246616",
      "stakeDisabled": false,
      "vault": null
    },
    {
      "id": "creativity-farm",
      "displayName": "Creativity Farm",
      "logo": "🎨",
      "tokenForLogo": "",
      "type": 0,
      "contractAddress": "0x59258F4e15A5fC74A7284055A8094F58108dbD4f",
      "autoStakePoolAddress": "",
      "collateralAddress": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardAPY": "0.00",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["FARM"],
      "detailsTokenSymbol": "Creativity",
      "lpTokenData": {
        "address": "0xa0246c9032bC3A600820415aE600c6388619A14D",
        "decimals": "18",
        "symbol": "FARM",
        "price": 110.54
      },
      "rewardAPR": "0",
      "rewardPerToken": "264895390986095530",
      "totalSupply": "2303841481353631355870",
      "finishTime": "1601563478",
      "totalValueLocked": "254666.6373488304100778698",
      "stakeDisabled": false,
      "vault": null
    },
    {
      "id": "weth-usdt-uni-farm",
      "displayName": "Uni ETH-USDT",
      "logo": "",
      "tokenForLogo": "ETH-USDT",
      "type": 0,
      "contractAddress": "0x75071F2653fBC902EBaff908d4c68712a5d1C960",
      "autoStakePoolAddress": "",
      "collateralAddress": "0x7ddc3fff0612e75ea5ddc0d6bd4e268f70362cff",
      "rewardAPY": "22.04",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["fUNI_LP_WETH_USDT"],
      "detailsTokenSymbol": "fUNISWAP_LP",
      "lpTokenData": {
        "address": "0x7DDc3ffF0612E75Ea5ddC0d6Bd4e268f70362Cff",
        "decimals": "18",
        "symbol": "fUNI-V2",
        "price": "49915731.60413436545238271767"
      },
      "rewardAPR": "19.958275516794586562",
      "rewardPerToken": "2627789175107516081421",
      "totalSupply": "735338280508354270",
      "finishTime": "1603828471",
      "totalValueLocked": "36772108.0884701346397125958",
      "stakeDisabled": false,
      "vault": {
        "logoUrl": "./icons/eth_usdt.svg",
        "apyIconUrl": "./icons/uni.png",
        "tokenAddress": "0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852",
        "decimals": "18",
        "vaultAddress": "0x7ddc3fff0612e75ea5ddc0d6bd4e268f70362cff",
        "cmcRewardTokenSymbols": ["FARM", "ETH", "USDT"],
        "pricePerFullShare": "1003877225162379827",
        "estimatedApy": "13.78",
        "underlyingBalanceWithInvestment": "738635266929820332",
        "usdPrice": "49813923.70308837487437353461"
      }
    },
    {
      "id": "weth-usdc-uni-farm",
      "displayName": "Uni ETH-USDC",
      "logo": "",
      "tokenForLogo": "ETH-USDC",
      "type": 0,
      "contractAddress": "0x156733b89Ac5C704F3217FEe2949A9D4A73764b5",
      "autoStakePoolAddress": "",
      "collateralAddress": "0xA79a083FDD87F73c2f983c5551EC974685D6bb36",
      "rewardAPY": "12.77",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["fUNI_LP_WETH_USDC"],
      "detailsTokenSymbol": "fUNISWAP_LP",
      "lpTokenData": {
        "address": "0xA79a083FDD87F73c2f983c5551EC974685D6bb36",
        "decimals": "18",
        "symbol": "fUNI-V2",
        "price": "49806346.93759631932002870041"
      },
      "rewardAPR": "12.031529566470629832",
      "rewardPerToken": "1778311265575285367947",
      "totalSupply": "1436021780486143076",
      "finishTime": "1603828685",
      "totalValueLocked": "71653918.86895244696405823708",
      "stakeDisabled": false,
      "vault": {
        "logoUrl": "./icons/eth_usdc.svg",
        "apyIconUrl": "./icons/uni.png",
        "tokenAddress": "0xb4e16d0168e52d35cacd2c6185b44281ec28c9dc",
        "decimals": "18",
        "vaultAddress": "0xA79a083FDD87F73c2f983c5551EC974685D6bb36",
        "cmcRewardTokenSymbols": ["FARM", "ETH", "USDC"],
        "pricePerFullShare": "1003345446719379630",
        "estimatedApy": "9.91",
        "underlyingBalanceWithInvestment": "1506041667117510032",
        "usdPrice": "49731142.48634137852474755732"
      }
    },
    {
      "id": "weth-dai-uni-farm",
      "displayName": "Uni ETH-DAI",
      "logo": "",
      "tokenForLogo": "UNI:ETH-DAI",
      "type": 0,
      "contractAddress": "0x7aeb36e22e60397098C2a5C51f0A5fB06e7b859c",
      "autoStakePoolAddress": "",
      "collateralAddress": "0x307E2752e8b8a9C29005001Be66B1c012CA9CDB7",
      "rewardAPY": "26.87",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["fUNI_LP_WETH_DAI"],
      "detailsTokenSymbol": "fUNISWAP_LP",
      "lpTokenData": {
        "address": "0x307E2752e8b8a9C29005001Be66B1c012CA9CDB7",
        "decimals": "18",
        "symbol": "fUNI-V2",
        "price": "47.33919201467446263138"
      },
      "rewardAPR": "23.855527086853381001",
      "rewardPerToken": "1872820305120187",
      "totalSupply": "672264937604022036780684",
      "finishTime": "1603828373",
      "totalValueLocked": "31824478.9659699458766490905",
      "stakeDisabled": false,
      "vault": null
    },
    {
      "id": "weth-wbtc-uni-farm",
      "displayName": "Uni ETH-WBTC",
      "logo": "",
      "tokenForLogo": "UNI:ETH-WBTC",
      "type": 0,
      "contractAddress": "0xF1181A71CC331958AE2cA2aAD0784Acfc436CB93",
      "autoStakePoolAddress": "",
      "collateralAddress": "0x01112a60f427205dcA6E229425306923c3Cc2073",
      "rewardAPY": "16.58",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["fUNI_LP_WETH_WBTC"],
      "detailsTokenSymbol": "fUNISWAP_LP",
      "lpTokenData": {
        "address": "0x01112a60f427205dcA6E229425306923c3Cc2073",
        "decimals": "18",
        "symbol": "fUNI-V2",
        "price": "494254401.09013042093254144188"
      },
      "rewardAPR": "15.366925119541846072",
      "rewardPerToken": "16814898105857063431869",
      "totalSupply": "105764392828010627",
      "finishTime": "1603828440",
      "totalValueLocked": "52370242.71705016368065711315",
      "stakeDisabled": false,
      "vault": null
    },
    {
      "id": "fweth-farm",
      "displayName": "WETH Farm",
      "logo": "",
      "tokenForLogo": "WETH",
      "type": 0,
      "contractAddress": "0x3DA9D911301f8144bdF5c3c67886e5373DCdff8e",
      "autoStakePoolAddress": "",
      "collateralAddress": "0xFE09e53A81Fe2808bc493ea64319109B5bAa573e",
      "rewardAPY": "10.13",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["fWETH"],
      "detailsTokenSymbol": "",
      "lpTokenData": {
        "address": "0xFE09e53A81Fe2808bc493ea64319109B5bAa573e",
        "decimals": "18",
        "symbol": "fWETH",
        "price": "393.00816223466192949532"
      },
      "rewardAPR": "9.660225293664012307",
      "rewardPerToken": "2435864046315467",
      "totalSupply": "21780451579541740897624",
      "finishTime": "1603836228",
      "totalValueLocked": "8559895.24791673918539009354",
      "stakeDisabled": false,
      "vault": {
        "logoUrl": "./icons/eth.png",
        "apyIconUrl": "",
        "tokenAddress": "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
        "decimals": "18",
        "vaultAddress": "0xFE09e53A81Fe2808bc493ea64319109B5bAa573e",
        "cmcRewardTokenSymbols": ["FARM", "WETH"],
        "pricePerFullShare": "1000504473497777372",
        "estimatedApy": "0.87",
        "underlyingBalanceWithInvestment": "22377071650901086671140",
        "usdPrice": "392.81"
      }
    },
    {
      "id": "fweth-sushi-wbtc-tbtc",
      "displayName": "SUSHI: WBTC-TBTC",
      "logo": "",
      "tokenForLogo": "SUSHI",
      "type": 0,
      "contractAddress": "0x9523FdC055F503F73FF40D7F66850F409D80EF34",
      "autoStakePoolAddress": "",
      "collateralAddress": "0xF553E1f826f42716cDFe02bde5ee76b2a52fc7EB",
      "rewardAPY": "11.44",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["fSUSHI-WBTC-TBTC"],
      "detailsTokenSymbol": "",
      "lpTokenData": {
        "address": "0xF553E1f826f42716cDFe02bde5ee76b2a52fc7EB",
        "decimals": "18",
        "symbol": "fSLP",
        "price": "2613086507.64519118776157308514"
      },
      "rewardAPR": "10.841679915597138404",
      "rewardPerToken": "75522504772009472359800",
      "totalSupply": "1525224082814143",
      "finishTime": "1603836423",
      "totalValueLocked": "3985542.47193714879966191383",
      "stakeDisabled": false,
      "vault": null
    },
    {
      "id": "farm-usdc",
      "displayName": "USDC Farm",
      "logo": "",
      "tokenForLogo": "USDC",
      "type": 0,
      "contractAddress": "0x4F7c28cCb0F1Dbd1388209C67eEc234273C878Bd",
      "autoStakePoolAddress": "",
      "collateralAddress": "0xf0358e8c3CD5Fa238a29301d0bEa3D63A17bEdBE",
      "rewardAPY": "23.03",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["fUSDC"],
      "detailsTokenSymbol": "",
      "lpTokenData": {
        "address": "0xf0358e8c3CD5Fa238a29301d0bEa3D63A17bEdBE",
        "decimals": "6",
        "symbol": "fUSDC",
        "price": "0.834979955199"
      },
      "rewardAPR": "20.764625827930042945",
      "rewardPerToken": "10390209949651724965753061",
      "totalSupply": "33385169506789",
      "finishTime": "1603836049",
      "totalValueLocked": "27875947.339089700146346011",
      "stakeDisabled": false,
      "vault": {
        "logoUrl": "./icons/usdc.png",
        "apyIconUrl": "./icons/curve.png",
        "tokenAddress": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
        "decimals": "6",
        "vaultAddress": "0xf0358e8c3CD5Fa238a29301d0bEa3D63A17bEdBE",
        "cmcRewardTokenSymbols": ["FARM", "USDC"],
        "pricePerFullShare": "835081",
        "estimatedApy": "0.00",
        "underlyingBalanceWithInvestment": "31501296002815",
        "usdPrice": "0.999879"
      }
    },
    {
      "id": "farm-dai",
      "displayName": "DAI Farm",
      "logo": "",
      "tokenForLogo": "DAI",
      "type": 0,
      "contractAddress": "0x15d3A64B2d5ab9E152F16593Cdebc4bB165B5B4A",
      "autoStakePoolAddress": "",
      "collateralAddress": "0xab7fa2b2985bccfc13c6d86b1d5a17486ab1e04c",
      "rewardAPY": "26.87",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["fDAI"],
      "detailsTokenSymbol": "",
      "lpTokenData": {
        "address": "0xab7FA2B2985BCcfC13c6D86b1D5A17486ab1e04C",
        "decimals": "18",
        "symbol": "fDAI",
        "price": "0.97979357088240902538"
      },
      "rewardAPR": "23.854099000538426417",
      "rewardPerToken": "10308468853523",
      "totalSupply": "1769001114949416144744299",
      "finishTime": "1603836228",
      "totalValueLocked": "1733255.9193112513636133011",
      "stakeDisabled": false,
      "vault": {
        "logoUrl": "./icons/dai.svg",
        "apyIconUrl": "./icons/curve.png",
        "tokenAddress": "0x6B175474E89094C44Da98b954EedeAC495271d0F",
        "decimals": "18",
        "vaultAddress": "0xab7fa2b2985bccfc13c6d86b1d5a17486ab1e04c",
        "cmcRewardTokenSymbols": ["FARM", "DAI"],
        "pricePerFullShare": "970092644438028738",
        "estimatedApy": "0.00",
        "underlyingBalanceWithInvestment": "1827583947981401895216069",
        "usdPrice": "1.01"
      }
    },
    {
      "id": "farm-usdt",
      "displayName": "USDT Farm",
      "logo": "",
      "tokenForLogo": "USDT",
      "type": 0,
      "contractAddress": "0x6ac4a7ab91e6fd098e13b7d347c6d4d1494994a2",
      "autoStakePoolAddress": "",
      "collateralAddress": "0x053c80eA73Dc6941F518a68E2FC52Ac45BDE7c9C",
      "rewardAPY": "26.09",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["fUSDT"],
      "detailsTokenSymbol": "",
      "lpTokenData": {
        "address": "0x053c80eA73Dc6941F518a68E2FC52Ac45BDE7c9C",
        "decimals": "6",
        "symbol": "fUSDT",
        "price": "0.844728"
      },
      "rewardAPR": "23.235051511290753712",
      "rewardPerToken": "12970215096029893525307272",
      "totalSupply": "29592633549668",
      "finishTime": "1603836228",
      "totalValueLocked": "24997726.153143950304",
      "stakeDisabled": false,
      "vault": {
        "logoUrl": "./icons/usdt.png",
        "apyIconUrl": "./icons/curve.png",
        "tokenAddress": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
        "decimals": "6",
        "vaultAddress": "0x053c80eA73Dc6941F518a68E2FC52Ac45BDE7c9C",
        "cmcRewardTokenSymbols": ["FARM", "USDT"],
        "pricePerFullShare": "844728",
        "estimatedApy": "0.00",
        "underlyingBalanceWithInvestment": "26332927868926",
        "usdPrice": "1"
      }
    },
    {
      "id": "farm-tusd",
      "displayName": "TUSD Farm",
      "logo": "",
      "tokenForLogo": "TUSD",
      "type": 0,
      "contractAddress": "0xeC56a21CF0D7FeB93C25587C12bFfe094aa0eCdA",
      "autoStakePoolAddress": "",
      "collateralAddress": "0x7674622c63Bee7F46E86a4A5A18976693D54441b",
      "rewardAPY": "9.09",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["fTUSD"],
      "detailsTokenSymbol": "",
      "lpTokenData": {
        "address": "0x7674622c63Bee7F46E86a4A5A18976693D54441b",
        "decimals": "18",
        "symbol": "fTUSD",
        "price": "1.000739647583816001"
      },
      "rewardAPR": "8.710139667625318857",
      "rewardPerToken": "54396290991778",
      "totalSupply": "19921809122980517147240257",
      "finishTime": "1603828191",
      "totalValueLocked": "19936544.24096357325267046454",
      "stakeDisabled": false,
      "vault": {
        "logoUrl": "./icons/tusd.png",
        "apyIconUrl": "./icons/tusd.png",
        "tokenAddress": "0x0000000000085d4780B73119b644AE5ecd22b376",
        "decimals": "18",
        "vaultAddress": "0x7674622c63Bee7F46E86a4A5A18976693D54441b",
        "cmcRewardTokenSymbols": ["FARM", "TUSD"],
        "pricePerFullShare": "1000739647583816001",
        "estimatedApy": "0.00",
        "underlyingBalanceWithInvestment": "19936552486223340122586468",
        "usdPrice": "1"
      }
    },
    {
      "id": "farm-wbtc",
      "displayName": "WBTC Farm",
      "logo": "",
      "tokenForLogo": "WBTC",
      "type": 0,
      "contractAddress": "0x917d6480Ec60cBddd6CbD0C8EA317Bcc709EA77B",
      "autoStakePoolAddress": "",
      "collateralAddress": "0x5d9d25c7C457dD82fc8668FFC6B9746b674d4EcB",
      "rewardAPY": "32.39",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["fWBTC"],
      "detailsTokenSymbol": "",
      "lpTokenData": {
        "address": "0x5d9d25c7C457dD82fc8668FFC6B9746b674d4EcB",
        "decimals": "8",
        "symbol": "fWBTC",
        "price": "13079.3787094738"
      },
      "rewardAPR": "28.131891379741221308",
      "rewardPerToken": "1443832472677162510143089284",
      "totalSupply": "93837889471",
      "finishTime": "1603836228",
      "totalValueLocked": "12286637.767429178021916462",
      "stakeDisabled": false,
      "vault": {
        "logoUrl": "./icons/wbtc.png",
        "apyIconUrl": "./icons/curve.png",
        "tokenAddress": "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599",
        "decimals": "8",
        "vaultAddress": "0x5d9d25c7C457dD82fc8668FFC6B9746b674d4EcB",
        "cmcRewardTokenSymbols": ["FARM", "WBTC"],
        "pricePerFullShare": "100094273",
        "estimatedApy": "0.00",
        "underlyingBalanceWithInvestment": "94800351533",
        "usdPrice": "13067.06"
      }
    },
    {
      "id": "farm-renbtc",
      "displayName": "RENBTC Farm",
      "logo": "",
      "tokenForLogo": "RENBTC",
      "type": 0,
      "contractAddress": "0x7b8Ff8884590f44e10Ea8105730fe637Ce0cb4F6",
      "autoStakePoolAddress": "",
      "collateralAddress": "0xC391d1b08c1403313B0c28D47202DFDA015633C4",
      "rewardAPY": "16.72",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["fRENBTC"],
      "detailsTokenSymbol": "",
      "lpTokenData": {
        "address": "0xC391d1b08c1403313B0c28D47202DFDA015633C4",
        "decimals": "8",
        "symbol": "frenBTC",
        "price": "13079.260627571"
      },
      "rewardAPR": "15.483556952433239885",
      "rewardPerToken": "1168326604450553148591643104",
      "totalSupply": "91188643825",
      "finishTime": "1603836228",
      "totalValueLocked": "11962803.999401413797058525",
      "stakeDisabled": false,
      "vault": {
        "logoUrl": "./icons/ren.png",
        "apyIconUrl": "./icons/curve.png",
        "tokenAddress": "0xeb4c2781e4eba804ce9a9803c67d0893436bb27d",
        "decimals": "8",
        "vaultAddress": "0xC391d1b08c1403313B0c28D47202DFDA015633C4",
        "cmcRewardTokenSymbols": ["FARM", "RENBTC"],
        "pricePerFullShare": "100031821",
        "estimatedApy": "0.00",
        "underlyingBalanceWithInvestment": "92462814845",
        "usdPrice": "13114.57"
      }
    },
    {
      "id": "crv-renbtc",
      "displayName": "CRVRENBTC Farm",
      "logo": "",
      "tokenForLogo": "CRVRENWBTC",
      "type": 0,
      "contractAddress": "0xA3Cf8D1CEe996253FAD1F8e3d68BDCba7B3A3Db5",
      "autoStakePoolAddress": "",
      "collateralAddress": "0x9aA8F427A17d6B0d91B6262989EdC7D45d6aEdf8",
      "rewardAPY": "2.22",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["fCRVRENWBTC"],
      "detailsTokenSymbol": "",
      "lpTokenData": {
        "address": "0x9aA8F427A17d6B0d91B6262989EdC7D45d6aEdf8",
        "decimals": "18",
        "symbol": "fcrvRenWBTC",
        "price": "13265.22765776781645767503"
      },
      "rewardAPR": "2.194331001607952616",
      "rewardPerToken": "43682333039218407",
      "totalSupply": "7329233600494738929804",
      "finishTime": "1603836228",
      "totalValueLocked": "97223952.26752400591523527884",
      "stakeDisabled": false,
      "vault": {
        "logoUrl": "./icons/crvrenwbtc.png",
        "apyIconUrl": "./icons/curve.png",
        "tokenAddress": "0x49849C98ae39Fff122806C06791Fa73784FB3675",
        "decimals": "18",
        "vaultAddress": "0x9aA8F427A17d6B0d91B6262989EdC7D45d6aEdf8",
        "cmcRewardTokenSymbols": ["FARM", "RENBTC", "WBTC"],
        "pricePerFullShare": "1000218487834929191",
        "estimatedApy": "0.00",
        "underlyingBalanceWithInvestment": "7331863335283303539098",
        "usdPrice": "13067.06"
      }
    }
  ],
  "Type1": [
    {
      "id": "profit-sharing-farm",
      "displayName": "Profit Sharing",
      "logo": "🚜",
      "tokenForLogo": "",
      "type": 1,
      "contractAddress": "0x8f5adc58b32d4e5ca02eac0e293d35855999436c",
      "autoStakePoolAddress": "0x25550Cccbd68533Fa04bFD3e3AC4D09f9e00Fc50",
      "collateralAddress": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardAPY": "450.79",
      "rewardToken": "0xa0246c9032bC3A600820415aE600c6388619A14D",
      "rewardTokenSymbols": ["FARM"],
      "collateralTokenSymbols": ["FARM"],
      "detailsTokenSymbol": "",
      "lpTokenData": {
        "address": "0xa0246c9032bC3A600820415aE600c6388619A14D",
        "decimals": "18",
        "symbol": "FARM",
        "price": 110.54
      },
      "rewardAPR": "",
      "rewardPerToken": "91993777172088914",
      "totalSupply": "141602904949727958326031",
      "finishTime": "1603844873",
      "totalValueLocked": "15652785.11314292851335946674",
      "stakeDisabled": false,
      "vault": null
    }
  ]
}
```

#### POST `/staking/harvestfarm/deposit_eth`

```shell
curl -d '{"user":"0x92E73408801e713f8371f8A8c31a40130ae61a40", "vault":"0x99b0d6641A63Ce173E6EB063b3d3AED9A35Cf9bf", "value":"1"}' -H 'Content-Type: application/json' http://localhost:9011/staking/harvestfarm/deposit_eth
```

Response:

```json
{
  "data": "0x03f9c79300000000000000000000000099b0d6641a63ce173e6eb063b3d3aed9a35cf9bf",
  "tx_fee": "",
  "value": "1000000000000000000",
  "contract_addr": "0x1c85CD673871EafEa31C0EdB42f42751aE64B5Bd"
}
```

value 的值是要传入的 ETH 的数量
