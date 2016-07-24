package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Allcurrencies struct {
	List struct {
		Meta struct {
			Type  string `json:"type"`
			Start int    `json:"start"`
			Count int    `json:"count"`
		} `json:"meta"`
		Resources []struct {
			Resource struct {
				Classname string `json:"classname"`
				Fields    struct {
					Name    string `json:"name"`
					Price   string `json:"price"`
					Symbol  string `json:"symbol"`
					Ts      string `json:"ts"`
					Type    string `json:"type"`
					Utctime string `json:"utctime"`
					Volume  string `json:"volume"`
				} `json:"fields"`
			} `json:"resource"`
		} `json:"resources"`
	} `json:"list"`
}

func main() {
	fmt.Println("EXCHANGE RATE Doller & British Pound")

	url := "http://finance.yahoo.com/webservice/v1/symbols/allcurrencies/quote?format=json"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var p Allcurrencies
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

	for i := 0; i < p.List.Meta.Count; i++ {
		if p.List.Resources[i].Resource.Fields.Name == "USD/GBP" {
			fmt.Println(p.List.Resources[i].Resource.Fields.Name)
			fmt.Println(p.List.Resources[i].Resource.Fields.Price)
			fmt.Println("GBP/USD")
			f, err := strconv.ParseFloat(p.List.Resources[i].Resource.Fields.Price, 64)
			if err != nil {
				panic(err)
			}
			f = 1 / f
			fmt.Printf("%.6f\n", f)

		}
	}
}
