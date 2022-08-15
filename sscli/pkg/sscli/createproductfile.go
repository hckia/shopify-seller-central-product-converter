package sscli

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/hckia/shopify-seller-central-product-converter/shopify-seller-api/pkg/swagger/server/models"
)

func createProductFile(makeParam string, respString []byte) bool {

	var responseObjects []*models.ProductRow
	json.Unmarshal(respString, &responseObjects)

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

	f, err := os.Create(makeParamStrFile)

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	defer f.Close()

	// Create the writer with the file
	writer := csv.NewWriter(f)
	defer writer.Flush()

	csvRows := make([][]string, len(responseObjects))

	for i, swatch := range responseObjects {
		csvRows[i] = []string{
			swatch.Handle,
			strconv.FormatInt(swatch.Year, 10),
			swatch.OptionName,
			swatch.OptionValue,
			strconv.FormatFloat(swatch.Price, 'f', 2, 32),
		}
	}
	err = writer.WriteAll(csvRows) // calls Flush internally

	if err != nil {
		log.Fatal("Cannot write to file...", err)
		return false
	} else {
		return true
	}
}

//	product.Handle = "honda-accord-2000-touch-up-kit"
//	product.OptionName = "honda accord Paint Colors"
//	product.OptionValue = "Dark Emerald Pearl : G-87P"
//	product.Price = 39.21
//	product.Year = 2000
