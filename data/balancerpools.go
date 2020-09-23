package data

func storeBalancerPools() {

	// var aBalancerPools types.BalancerPools

	// f, err := ioutil.ReadFile("./resources/balancer.json")

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// if err := json.Unmarshal(f, &aBalancerPools); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// for _, j := range aBalancerPools.Data.Pools {
	// 	for _, aSwap := range j.Swaps {
	// 		poolKey := fmt.Sprintf("%s-%s", aSwap.TokenInSym, aSwap.TokenOutSym)

	// 		BalancerPools[poolKey] = j.ID
	// 	}
	// }

	// WETH - DAI
	BalancerPools["WETH-DAI"] = "0x99e582374015c1d2f3c0f98d0763b4b1145772b7"
	BalancerPools["DAI-WETH"] = "0x99e582374015c1d2f3c0f98d0763b4b1145772b7"

	// WETH - USDC
	BalancerPools["WETH-USDC"] = "0xe969991ce475bcf817e01e1aad4687da7e1d6f83"
	BalancerPools["USDC-WETH"] = "0xe969991ce475bcf817e01e1aad4687da7e1d6f83"

	// WETH - YFI
	BalancerPools["YFI-WETH"] = "0x41284a88D970D3552A26FaE680692ED40B34010C"
	BalancerPools["WETH-YFI"] = "0x41284a88D970D3552A26FaE680692ED40B34010C"

	// USDC - DAI
	BalancerPools["DAI-USDC"] = "0x9B208194Acc0a8cCB2A8dcafEACfbB7dCc093F81"
	BalancerPools["USDC-DAI"] = "0x9B208194Acc0a8cCB2A8dcafEACfbB7dCc093F81"

	// DAI - SAFE
	BalancerPools["SAFE-DAI"] = "0x3031745e732dce8fecccc94aca13d5fa18f1012d"
	BalancerPools["DAI-SAFE"] = "0x3031745e732dce8fecccc94aca13d5fa18f1012d"

	// DAI - YFII
	BalancerPools["DAI-YFII"] = "0x16cAC1403377978644e78769Daa49d8f6B6CF565"
	BalancerPools["YFII-DAI"] = "0x16cAC1403377978644e78769Daa49d8f6B6CF565"

	// DAI - BAL
	BalancerPools["BAL-DAI"] = "0x2E41132dab88A9bAd80740A1392D322Bf023d494"
	BalancerPools["DAI-BAL"] = "0x2E41132dab88A9bAd80740A1392D322Bf023d494"
}
