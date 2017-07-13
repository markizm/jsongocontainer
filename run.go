package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Formatted to match layout on dev-catchpoint../api/tests

type Xjson struct {
	Data []struct {
		Attributes struct {
			Name     string "json:name"
			Test_url string "json:test_url"
		}
	}
}

func getJson(w http.ResponseWriter, r *http.Request) {
	url := "http://dev-catchpoint.ssvm.com:8080/api/tests"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("unable to get data")
	}
	defer res.Body.Close()

	obj, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("sum ting wong")
	}

	var j Xjson

	err = json.Unmarshal(obj, &j)
	if err != nil {
		fmt.Println("mo wong gur")
	}

	json.NewEncoder(w).Encode(j)

}

func main() {

	http.HandleFunc("/get", getJson)
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.ListenAndServe(":9085", nil)
}
