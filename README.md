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

GET: `tokenlist`

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


## Contract

### UniswapV2


POST `/swapInfo`

contract: `UniswapV2`
from: `ETH`
to: `DAI`
amount: `10000000000000000` // 0.01 ETH
user: `0xcd69c8CbBFe5b1219C0f8911204aA961294E74e3`
slippage: `5` // 0.05% 的滑点

```json
{
  "data": "0x7ff36ab5000000000000000000000000000000000000000000000000002382664887b0000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000cd69c8cbbfe5b1219c0f8911204aa961294e74e3000000000000000000000000000000000000000000000000000000005f6083900000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000006b175474e89094c44da98b954eedeac495271d0f",
  "tx_fee": "",
  "contract_addr": "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",
  "from_token_amount": "",
  "to_token_amount": ""
}
```

UniswapV2接口说明：
+ Addr: `0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D`

+ 方法说明：

+ `7ff36ab5` swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline)

> 把一定量的ETH转为某Token
>
> `amountOutMin`:  可以传一个比较小的数字 比如 1000000 (这个如何设置待查询)
>
> `path`: [address(0), address(Token)];
>
> `to`: 用户自己
>
> `deadline`: 当前时间戳+60秒

+ `38ed1739` swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)

> 把一定量的Token转为另一个Token
>
> 参数同上

+ `18cbafe5` swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)

> 把一定量的Token转为ETH
>
> 参数同上



下面暂时用不到：

+ `791ac947` swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)
+ `fb3bdb41` swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline)
+ `4a25d94a` swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline)



### Bancor接口说明

+ Addr: `0x2F9EC37d6CcFFf1caB21733BdaDEdE11c823cCB0`

+ `0xe57738e5` claimAndConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee)
+ `0x569706eb` convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee)
+ `0xb77d239b` convertByPath(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee)


### Mooniswap接口说明

+ 首先查询Factory位置： 通过Addr: `0x71CD6666064C3A1354a3B4dca5fA1E2D3ee7D303` 的 `pools` 方法

> pools(address from, address to)

+ 然后调用Factory的方法
+ `0xd5bcb9b5` swap(address src, address dst, uint256 amount, uint256 minReturn, address referral)

> 参数说明:
>
> `src`: from Token
>
> `dst`: to Token
>
> `minReturn`: 10000 (待确定)
>
> `referral`: `0x68a17b587caf4f9329f0e372e3a78d23a46de6b5` （这是Mooniswap Factory合约的Owner）



### Kyber接口说明

+ Addr: `0x818E6FECD516Ecc3849DAf6845e3EC868087B755`
+ `0xcb3c28c7` trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId)

> 参数说明：
>
> `src`: FromToken Addr
>
> `srcAmount`: FromToken Amount
>
> `dest`: ToToken Addr  (如果是ETH，则是`0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee`)
>
> `destAddress`: 用户地址
>
> `maxDestAmount`:
>
> `minConversionRate`:
>
> `walletId`:  `0xf1aa99c69715f423086008eb9d06dc1e35cc504d`

+ `0x29589f61` tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint)

### 1inch说明

+ Addr: `0xC586BeF4a0992C495Cf22e1aeEE4E446CECDee0E`
+ `0xe2a7515e` swap(address fromToken, address toToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 featureFlags)

### SushiSwap
+ Addr: `0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F`

+ `0x18cbafe5` swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)
+ `0x7ff36ab5` swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline)
+ `0x38ed1739` swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)

### Dforce
+ Addr: `0x03eF3f37856bD08eb47E2dE7ABc4Ddd2c19B60F2`
+ `0xdf791e50` swap(address source, address dest, uint256 sourceAmount)

### Paraswap接口说明

+ Addr：

### ZeroX接口说明

+ Addr






```
// API Call：
// 0x
// tokenlan TODO:change API

// Contract Call：
// bancor
// Curve
// dforce
// kyber
// mooniswap
// oasis
// 1inch
// Sushiswap
// UniswapV2
// Paraswap

// TODO
// dYdX
// 0x
// balancer
// DDEX
// Loopring
// DoDo
// IDEX
// DEX.AG
// Tokenlon
```

+ `tokenAddr.json` is from: `https://api.paraswap.io/v2/tokens`