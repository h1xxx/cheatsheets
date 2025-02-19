package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type infoT struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	City    string `json:"city"`
}

func main() {
	req, err := http.NewRequest("GET", "https://am.i.mullvad.net/json", nil)
	req.Header.Set("User-Agent", "curl/7.74.0")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var info infoT

	err = json.Unmarshal(body, &info)
	if err != nil {
		panic(err)
	}

	out, err := json.MarshalIndent(info, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", out)
}
