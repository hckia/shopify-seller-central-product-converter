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

func createSwatchFile(makeParam string, respString []byte) bool {

	var responseObjects []*models.SwatchRow
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

	headerRow := []string{
		"Make",
		"Model",
		"Year",
		"MMY",
		"Color name",
		"Color code",
		"Variant Metafield: variants.color [color]",
		"Tricoat",
		"Handle",
	}
	tempRow := make([][]string, 1)

	tempRow[0] = headerRow
	csvRows = append(tempRow, csvRows...)

	err = writer.WriteAll(csvRows) // calls Flush internally

	if err != nil {
		log.Fatal("Cannot write to file...", err)
		return false
	} else {
		return true
	}

}
