### Example

POST `/handler`
+ from: `USDT`
+ to: `DAI`
+ amount: `12000`

```json
{
  "from_name": "USDT",
  "to_name": "DAI",
  "from_addr": "0xdac17f958d2ee523a2206206994597c13d831ec7",
  "to_addr": "0x6b175474e89094c44da98b954eedeac495271d0f",
  "exchange_pairs": [
    {
      "contract_name": "Paraswap",
      "ratio": "11739248500991466"
    },
    {
      "contract_name": "1inch",
      "ratio": "17838169486919069"
    },
    {
      "contract_name": "UniswapV2",
      "ratio": "11702340661500501"
    },
    {
      "contract_name": "Dforce",
      "ratio": "11673850165594800"
    },
    {
      "contract_name": "Kyber",
      "ratio": "973500888775927666"
    },
    {
      "contract_name": "Bancor",
      "ratio": "338627228055330043560"
    },
    {
      "contract_name": "ZeroX",
      "ratio": "0.974782"
    },
    {
      "contract_name": "Curve",
      "ratio": "11739248500991466"
    },
    {
      "contract_name": "Mooniswap",
      "ratio": "11673761524555519"
    }
  ]
}
```