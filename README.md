### Example

POST `/handler`
: ?from=USDC&to=DAI&amount=12`

```json
{
  "from_name": "USDC",
  "to_name": "DAI",
  "from_addr": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
  "to_addr": "0x6b175474e89094c44da98b954eedeac495271d0f",
  "exchange_pairs": [
    {
      "contract_name": "Paraswap",
      "ratio": ""
    },
    {
      "contract_name": "",
      "ratio": ""
    },
    {
      "contract_name": "ZeroX",
      "ratio": ""
    },
    {
      "contract_name": "Dforce",
      "ratio": "11668018489020"
    },
    {
      "contract_name": "UniswapV2",
      "ratio": "11640104980246"
    },
    {
      "contract_name": "1inch",
      "ratio": "5212621599959014"
    },
    {
      "contract_name": "Bancor",
      "ratio": ""
    },
    {
      "contract_name": "Mooniswap",
      "ratio": "11695176277650"
    },
    {
      "contract_name": "Curve",
      "ratio": "0"
    }
  ]
}
```