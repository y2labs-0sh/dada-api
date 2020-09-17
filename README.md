## Example

## POST `/aggrInfo`

+ from: `USDT`
+ to: `DAI`
+ amount: `1200000`

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

### Get: `tokenlist`

```
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

### POST: `swapInfo`

+ `contract`: `UniswapV2`
+ `from`: `ETH`
+ `to`: `DAI`
+ `amount`: `10000000`
+ `user`: `0xcd69c8CbBFe5b1219C0f8911204aA961294E74e3` (用户地址)
+ `slippage`: `5` (滑点为0.05%)

```
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

+ 当 `from` 为ETH时，不用去执行approve

+ 当`from`为其他ERC20 token时，如果allowance_satisfied 为 false，需要执行 approve，执行的合约地址为`from_token_addr`, call_data 为：`allowance_data`
+ `contract` 可以选择为：`UniswapV2` ，`Bancor`，`Dforce`，`Kyber`，`Mooniswap`，`Sushiswap`









