package sscli

import (
	"os"
	"log"
	"net/http"
	//"github.com/hckia/shopify-seller-central-product-converter/shopify-seller-api/pkg/swagger/server/models"
)

func fetchSwatches(make string, path string) {
	os.Setenv("ADD_SELLER_PATH", path)
	var URL = "localhost:8080/swatch/"+make

	request, err := http.NewRequest(
		http.MethodGet,
		URL,
		nil,
	)

	if err != nil {
		log.Println("Could not request make")
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "SSCLI (https://github.com/hckia/shopify-seller-central-product-converter/sscli)")

}
