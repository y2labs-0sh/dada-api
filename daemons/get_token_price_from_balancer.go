package daemons

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

func (d *tokenPriceBalancer) fetch() {
	resp, err := httpClient.Post(TokenPriceBalancerURI, "application/json", strings.NewReader(d.graphQL))
	if err != nil {
		d.logger.Error("Balancer Daemon: ", err.Error())
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		d.logger.Error("Balancer Daemon: ", err.Error())
		return
	}

	data_ := new(struct {
		Data struct {
			TokenPrices []PriceInfo `json:"tokenPrices"`
		} `json:"data"`
	})

	if err := json.Unmarshal(bodyBytes, data_); err != nil {
		d.logger.Error("Balancer Daemon: ", err.Error())
		return
	}

	d.listLock.Lock()
	d.list = data_.Data.TokenPrices
	d.listLock.Unlock()

	d.listLock.RLock()
	bs, err := json.Marshal(d.list)
	d.listLock.RUnlock()
	if err != nil {
		d.logger.Error("Balancer Daemon: ", err.Error())
		return
	}

	if err := d.fileStorage.save(bs); err != nil {
		d.logger.Error("Balancer Daemon: ", err.Error())
		return
	}
}
