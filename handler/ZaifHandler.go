package handler

import (
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
	"github.com/antonholmquist/jason"
)

type Price struct {
	json.Number `json:"price"`
}

func ZaifHandler(response http.ResponseWriter, _ *http.Request) {
	price := Price{GetLastPrice()}

	js, err := json.Marshal(price)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Write(js)
}


func GetLastPrice() json.Number {
	const url = "https://api.zaif.jp/api/1/last_price/btc_jpy"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	obj, err := jason.NewObjectFromBytes([]byte(body))

	var price json.Number
	price, err = obj.GetNumber("last_price")

	return price
}
