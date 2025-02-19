package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var PROXY = "http://localhost:8080"

func main() {
	proxyURL, err := url.Parse(PROXY)
	if err != nil {
		panic(err)
	}

	var PTransport = &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		//Proxy: http.ProxyFromEnvironment,   //https_proxy
	}

	timeout := time.Duration(5 * time.Second)

	client := http.Client{
		Transport: PTransport,
		Timeout:   timeout,
	}

	req, err := http.NewRequest("GET", "https://am.i.mullvad.net/json", nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBytes)
	fmt.Printf("GET Response = %s \n", string(bodyString))
}
