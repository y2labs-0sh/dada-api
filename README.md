### Example

POST `/handler`
+ from: `USDT`
+ to: `DAI`
+ amount: `1200000`

```json
{
  "from_name": "USDT",
  "to_name": "DAI",
  "from_addr": "0xdac17f958d2ee523a2206206994597c13d831ec7",
  "to_addr": "0x6b175474e89094c44da98b954eedeac495271d0f",
  "exchange_pairs": [
    {
      "contract_name": "Kyber",
      "ratio": "1164246077920575837600000000"
    },
    {
      "contract_name": "1inch",
      "ratio": "1172891899936796747739"
    },
    {
      "contract_name": "UniswapV2",
      "ratio": "1169381437382589109255"
    },
    {
      "contract_name": "ZeroX",
      "ratio": "1165302000"
    },
    {
      "contract_name": "Dforce",
      "ratio": "1164689295430342638085"
    },
    {
      "contract_name": "Bancor",
      "ratio": "311926868734838695747"
    },
    {
      "contract_name": "Mooniswap",
      "ratio": "1160003356690562127850"
    },
    {
      "contract_name": "Curve",
      "ratio": "51548921985186209306"
    },
    {
      "contract_name": "Paraswap",
      "ratio": "1172891900671635500000"
    }
  ]
}
```
