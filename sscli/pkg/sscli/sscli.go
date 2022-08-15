package sscli

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/hckia/shopify-seller-central-product-converter/shopify-seller-api/pkg/swagger/server/models"
)

func FetchSwatches(makeParam, path string) string {
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

	var responseObjects []*models.SwatchRow
	json.Unmarshal(respString, &responseObjects)

	for _, swatch := range responseObjects {
		fmt.Println(swatch.Handle)
	}

	var makeParamPath = os.Getenv("ADD_SELLER_PATH") + "/" + "Data/" + "Swatches/" + makeParam

	var makeParamStrFile string = makeParamPath + "/" + makeParam + ".csv"

	if _, err := os.Stat(makeParamStrFile); err == nil {
		log.Println("File already exists")
		os.Exit(0)
	} else if _, err := os.Stat(makeParamPath); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		err := os.MkdirAll(makeParamPath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	//fmt.Println("PATH: " + makeParamStrFile)

	f, err := os.Create(makeParamStrFile)
	defer f.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	// Create the writer with the file
	writer := csv.NewWriter(f)
	defer writer.Flush()

	csvRows := make([][]string, len(responseObjects))

	for i, swatch := range responseObjects {
		csvRows[i] = []string{
			swatch.Make,
			swatch.Model,
			strconv.FormatInt(swatch.Year, 10),
			swatch.Mmy,
			swatch.ColorName,
			swatch.ColorCode,
			swatch.HexCode,
			strconv.FormatBool(swatch.Tricoat),
			swatch.Handle,
		}
	}

	err = writer.WriteAll(csvRows) // calls Flush internally

	if err != nil {
		log.Fatal("Cannot write to file...", err)
	}
	// for i := 0; i < len(responseObjects); i++ {
	// 	fmt.Println(responseObjects[i])
	// }
	//fmt.Println(responseObject)

	// j, _ := json.Marshal(resp.Body)
	// jval := resp.Body
	// log.Println(respString)
	// log.Println("respString")
	// log.Print(resp.Body.Handle)
	// log.Println("resp.Body")
	// log.Print(j)
	// log.Println("j")
	// log.Println(jval)
	// log.Println("jval")
	var success = "SUCCESS"
	return success

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
