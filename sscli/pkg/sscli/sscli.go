package sscli

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func FetchSwatches(makeParam, path string) bool {
	os.Setenv("ADD_SELLER_PATH", path)
	var URL = "http://127.0.0.1:8080/swatch/" + makeParam

	request, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		log.Printf("Could not makeParam a request. %v", err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "SSCLI (https://github.com/hckia/shopify-seller-central-product-converter/sscli)")

	log.Println(URL)

	client := &http.Client{}

	resp, err := client.Do(request)

	if err != nil {
		log.Printf("Could not makeParam a request. %v", err)
	}

	defer resp.Body.Close()

	respString, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}
	fileResult := createSwatchFile(makeParam, respString)

	return fileResult

}

func FetchProducts(makeParam, path string) string {
	os.Setenv("ADD_SELLER_PATH", path)
	var URL = ":8080/product/" + makeParam

	request, err := http.NewRequest(
		http.MethodGet,
		URL,
		nil,
	)

	if err != nil {
		log.Println("Could not request makeParam")
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "SSCLI (https://github.com/hckia/shopify-seller-central-product-converter/sscli)")

	var success = "SUCCESS"
	return success

}
