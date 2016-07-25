package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// used https://mholt.github.io/json-to-go/
type beerYeast struct {
	YEASTS struct {
		YEAST []struct {
			NAME           string `json:"NAME"`
			TYPE           string `json:"TYPE"`
			FORM           string `json:"FORM"`
			AMOUNT         string `json:"AMOUNT"`
			LABORATORY     string `json:"LABORATORY"`
			PRODUCTID      string `json:"PRODUCT_ID"`
			MINTEMPERATURE string `json:"MIN_TEMPERATURE"`
			MAXTEMPERATURE string `json:"MAX_TEMPERATURE"`
			FLOCCULATION   string `json:"FLOCCULATION"`
			ATTENUATION    string `json:"ATTENUATION"`
			NOTES          string `json:"NOTES"`
			BESTFOR        string `json:"BEST_FOR"`
			CULTUREDATE    string `json:"CULTURE_DATE"`
		} `json:"YEAST"`
	} `json:"YEASTS"`
}

func main() {
	//location of MicroBeerServer
	url := "http://localhost:8080/yeast"
	//REST GET response at target URL
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed pull json URL: %s", err)
	}
	defer res.Body.Close() //idiomatic close with defer

	//convert to bytes
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ReadAll file failed: %s", err)
	}

	//unpack JSON response into struct
	var p beerYeast
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	//loop over return JSON object test print some values
	for _, yeast := range p.YEASTS.YEAST {
		fmt.Println(yeast.NAME)
		fmt.Println(yeast.LABORATORY)
		fmt.Println(yeast.BESTFOR)
		fmt.Println("")
	}

	// write the whole body at once to tmp file
	err = ioutil.WriteFile("outputTest.txt", body, 0644)
	if err != nil {
		log.Fatalf("writing JSON to File failed: %s", err)
	}

	fmt.Println("done")
}
