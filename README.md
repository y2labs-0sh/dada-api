### 例子1: 从1inch获取交易信息

POST: `http://127.0.0.1:9000/handler_1inch`

+ 参数：
  + `from`: ETH
  + `to`: DAI
  + `amount`: 2

结果：

```json
{
  "exchange_pair": [
    {
      "from_name": "ETH",
      "to_name": "DAI",
      "from_addr": "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
      "to_addr": "0x6b175474e89094c44da98b954eedeac495271d0f",
      "ratio": "125629494213209833123089",
      "tx_gas": ""
    }
  ]
}
```

### 例子2: 从bancor获取交易信息

POST: `127.0.0.1:9000/handler_bancor`

+ 参数：
  + `from`: ETH
  + `to`: ETH
  + `amount`: 2

结果：
```json
{
  "from_name": "ETH",
  "to_name": "DAI",
  "from_addr": "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
  "to_addr": "0x6b175474e89094c44da98b954eedeac495271d0f",
  "ratio": "867224646983205547030",
  "tx_gas": ""
}
```

例子3: 从paraswap获取交易信息：
POST: `127.0.0.1:9000/handler_paraswap`
+ 参数：
  + `from`: ETH
  + `to`: ETH
  + `amount`: 2
结果：
```json
{
  "from_name": "ETH",
  "to_name": "DAI",
  "from_addr": "",
  "to_addr": "",
  "ratio": "874015815707121700000",
  "tx_gas": ""
}
```
例子4: 从handlerKyberswap 获取交易信息:
POST: `http://127.0.0.1:9000/handler_kyberswap`
+ 参数：
  + `from`: ETH
  + `to`: ETH
  + `amount`: 2

结果：
```json
{
  "from_name": "ETH",
  "to_name": "DAI",
  "from_addr": "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
  "to_addr": "0x6b175474e89094c44da98b954eedeac495271d0f",
  "ratio": "430537615787260619137",
  "tx_gas": ""
}
```

例子5: 从handler_zeroX 获取交易信息:
POST： `http://127.0.0.1:9000/handler_zeroX`

+ 参数：
  + `from`: ETH
  + `to`: ETH
  + `amount`: 2

```json
{
  "from_name": "ETH",
  "to_name": "DAI",
  "from_addr": "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
  "to_addr": "0x6b175474e89094c44da98b954eedeac495271d0f",
  "ratio": "434.636027870297871685",
  "tx_gas": ""
}
```