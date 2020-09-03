package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// https://api.bancor.network/0.1/currencies/5937d635231e97001f744267/value?toCurrencyId=5e736e776ea615ea177bedf9&fromAmount=3000000000000000000&streamId=loadValue

func handlerBancor(c echo.Context) error {

	// TODO:1. 初始化放在单独一个函数中，只做一遍
	bancor := make(map[string]string)
	bancor["ETH"] = "5937d635231e97001f744267"
	bancor["DAI"] = "5e736e776ea615ea177bedf9"
	bancor["EOS"] = "5a1eb3753203d200012b8b75"
	bancor["BNT"] = "594bb7e468a95e00203b048d"
	bancor["BNB"] = "5a1fb3df634e000001878563"
	bancor["BAT"] = "5a1fd517634e000001878565"
	bancor["OMG"] = "5a086f93875e890001605abc"
	bancor["DICE"] = "5c0e8fdb675bcfda09467194"
	bancor["BLACK"] = "5c0e3f4464f6f90710095f3c"
	bancor["IQ"] = "5c0e62bb675bcfb491451108"
	bancor["MANA"] = "5a2cfacad0129700019a7270"
	bancor["SRN"] = "5a548d3244680b000198e7f3"
	bancor["LOC"] = "5b27eb823751b7bb8dad17bb"
	bancor["MKR"] = "5a72f33c6ce1040001bb5d56"
	bancor["ENJ"] = "5a174c5145a97200011ad30a"
	bancor["SNT"] = "5a2d307c47bbf500018ecc6e"
	bancor["GTO"] = "5a806bfda47fc50001d5bd5c"
	bancor["HORUS"] = "5c0e5181c01d8814fa2296f0"
	bancor["POWR"] = "5a12993ac1a3da0001a37aa6"
	bancor["EMCO"] = "5c84d2db49aae47399d738f2"
	bancor["KNC"] = "5a69a9cb1b67798dab40b997"
	bancor["XDCE"] = "5a60937501b4a8000153a22e"
	bancor["XPAT"] = "5b056686f2e6af420feecbd0"
	bancor["PLR"] = "5a69a9cf1b67798dab40b9ac"
	bancor["EPRA"] = "5c0e640cc01d8846c42327ad"
	bancor["AMN"] = "5ac5f1ceb9b20e2a231fc556"
	bancor["MEV"] = "5c0e698a48ded4568c33eea3"
	bancor["STX"] = "59d27d45acb3c12634d19efb"
	bancor["MEETONE"] = "5c0e529564f6f94efa0a075d"
	bancor["CEEK"] = "5b0fdbcd3b7891d727d11e4f"
	bancor["MYB"] = "5b164627ae2482321708eb93"
	bancor["TKN"] = "5a2d4ab9d0129700019a7604"
	bancor["X8X"] = "5a9bea027e33d66aea309c5d"
	bancor["SPD"] = "5b38b7d5e98d1bd535498422"
	bancor["XNK"] = "5b05322332e5aaa0a77c26c4"
	bancor["AGRI"] = "5bab52f508e2ef92e2100024"
	bancor["REAL"] = "5b27eb8d3751b7bb8dad17de"
	bancor["RBLX"] = "5af2e3e4b2007855dcaff39a"
	bancor["ZIPT"] = "5ae5bf1c3c124c8c3d024e02"
	bancor["CAN"] = "5a6f61ece3de16000123763a"
	bancor["WAND"] = "59f6f966a149b500018d7da4"
	bancor["DTRC"] = "5a3789b5a88c2a00013b61fc"
	bancor["OCT"] = "5c0e67d364f6f9d3670aa4c7"
	bancor["VEE"] = "5a5eeb3fc5e04600015b40f1"
	bancor["POA20"] = "5af31eb33a695919b3e7c2c3"
	bancor["REQ"] = "5a69a9cd1b67798dab40b9a1"
	bancor["MFT"] = "5b3e433ba5281a534e93166f"
	bancor["RVT"] = "5a69a9d51b67798dab40b9d7"
	bancor["ELF"] = "5a722b4cec75fc00012d4b8c"
	bancor["ANT"] = "5a37a5d7ed8a500001de3704"
	bancor["EURS"] = "5b795e4fafb461ec10127894"
	bancor["GNO"] = "59d745ff90509add31e9db14"
	bancor["AID"] = "5a7c2d3716bc390001ebfd35"
	bancor["TNS"] = "5ae816abc6295138b0a311d8"
	bancor["TAEL"] = "5c473ab671d550188b192ff7"
	bancor["GRID"] = "5a69a9cf1b67798dab40b9ad"
	bancor["ZINC"] = "5a8d5312fe85bd000167bc52"
	bancor["HVT"] = "5c0e662c54ed33261ddb853b"
	bancor["SVD"] = "5ab8b40552642aedbec8638f"
	bancor["MRG"] = "5afad1e8a40e8ccd21fa1a7f"
	bancor["EVO"] = "5bd023687449be2a298608a4"
	bancor["LOCI"] = "5a76c546ec75fc00012f05b6"
	bancor["BETR"] = "5a89a0538eae3b0001747e88"
	bancor["LDC"] = "5a97ec891cbcd07b03616d3b"
	bancor["MFG"] = "5aafdef1aa258beb09fd43a4"
	bancor["AUC"] = "5ad9c1d54c4998a2f940e933"
	bancor["LINK"] = "5b27eb593751b7bb8dad1735"
	bancor["DRT"] = "5a19bc169f604e00011f09ea"
	bancor["WLK"] = "5a33bfed9416220001fa71ce"
	bancor["RLC"] = "5a54b9cfb6b5870001b890e0"
	bancor["RCN"] = "5a92b443ef48d60001102970"

	// https: //api.bancor.network/0.1/currencies/5937d635231e97001f744267/value?toCurrencyId=5e736e776ea615ea177bedf9&fromAmount=9000000000000000000&streamId=loadValue

	baseURL := "https://api.bancor.network/0.1/currencies/%s/value?toCurrencyId=%s&fromAmount=%d&streamId=loadValue"
	from := c.FormValue("from")
	to := c.FormValue("to")
	amount := c.FormValue("amount")
	s, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "amount err: amount should be numeric")
	}

	queryURL := fmt.Sprintf(baseURL, bancor[from], bancor[to], int64(s*1000000000000000000)) // TODO:2.检查Bancor网站，有些Token基数不一样

	out := BancorResult{}
	resp, _ := http.Get(queryURL)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal([]byte(body), &out)

	result2 := exchange_pair{
		FromName: from,
		ToName:   to,
		FromAddr: m1[from].Address,
		ToAddr:   m1[to].Address,
		Ratio:    out.Data,
	}

	return c.JSON(http.StatusOK, result2)
}

type BancorResult struct {
	Data string `json:"data"`
}
