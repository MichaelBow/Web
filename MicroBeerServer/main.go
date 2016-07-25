package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/*a beer yeast microservice
* Simple HTTP server/client in golang reads json file from disk
* Client reads and stores and prints test data & test file
 */

func serveJSONFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get absolute path to json file from current directory
	pwd, _ := os.Getwd()
	//Try open json file
	jsonFile, err := os.Open(pwd + "/yeast.json")
	if err != nil {
		fmt.Println("Error opening json file:", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading json file:", err)
		return
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, jsonData, "", "\t")
	if err != nil {
		fmt.Println("Error making json human readable:", err)
		return
	}
	w.Write([]byte(prettyJSON.Bytes()))
}

func main() {
	fmt.Println("server start...")
	//create a http server
	mux := http.NewServeMux()
	mux.HandleFunc("/yeast", serveJSONFile)

	http.ListenAndServe(":8080", mux)

}
