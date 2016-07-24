package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Take the URL with your json data eg. http://finance.yahoo.com/webservice/v1/symbols/allcurrencies/quote?format=json
// convert to struct format at https://mholt.github.io/json-to-go/
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

// storing a currency as type string ready for interface
type Currency struct {
	price string
}

//function to convert to float then trim and convert back to string with .6 decimal precision
func (z Currency) inverse() string {
	price, err := strconv.ParseFloat(z.price, 64)
	if err != nil {
		panic(err)
	}
	//return as swapped currency value pair
	price = 1 / price
	return fmt.Sprintf("%.6f", price)
}

//interface to return a converted and inverse currency pair
type CurrencySwap interface {
	inverse() string
}

//method to format the float
func inverse(z CurrencySwap) {
	fmt.Println(z.inverse())
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
			//f, err := strconv.ParseFloat(p.List.Resources[i].Resource.Fields.Price, 64)
			//if err != nil {
			//	panic(err)
			//}
			//f = 1 / f
			//fmt.Printf("%.6f\n", f)
			//replace the above with a interface call
			money := Currency{p.List.Resources[i].Resource.Fields.Price}
			fmt.Println(money.inverse()) //value with attached method
			inverse(money) //interface returning std.o (println)

		}
	}

}
