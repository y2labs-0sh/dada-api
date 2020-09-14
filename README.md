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
UniswapV2接口说明：
+ Addr: `0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D`

+ `7ff36ab5` swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline)
+ `791ac947` swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)
+ `fb3bdb41` swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline)
+ `38ed1739` swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)
+ `4a25d94a` swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline)
+ `18cbafe5` swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)

### Bancor接口说明
+ Addr: `0x2F9EC37d6CcFFf1caB21733BdaDEdE11c823cCB0`

+ `0xe57738e5` claimAndConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee)
+ `0x569706eb` convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee)
+ `0xb77d239b` convertByPath(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee)


### Mooniswap接口说明



### Kyber接口说明

+ Addr: `0x818E6FECD516Ecc3849DAf6845e3EC868087B755`

+ `0xcb3c28c7` trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId)

+ `0x29589f61` tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint)


### Paraswap接口说明


### ZeroX接口说明

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
