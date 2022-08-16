package sscli

import (
	//"encoding/csv" - will be applied in future iterations as a flag.
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/hckia/shopify-seller-central-product-converter/shopify-seller-api/pkg/swagger/server/models"
)

func createProductFile(makeParam string, respString []byte) bool {

	var responseObjects []*models.ProductRow

	json.Unmarshal(respString, &responseObjects)

	var makeParamPath = os.Getenv("ADD_SELLER_PATH") + "/" + "Data/" + "Products/" + makeParam

	var makeParamStrFile string = makeParamPath + "/" + makeParam + ".xlsx"

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

	// EXCEL
	mf := excelize.NewFile()
	mi := mf.NewSheet(makeParam)

	mf.SetCellValue(makeParam, "A1", "Handle")
	mf.SetCellValue(makeParam, "B1", "year")
	mf.SetCellValue(makeParam, "C1", "Option1 Name")
	mf.SetCellValue(makeParam, "D1", "Option1 Value")
	mf.SetCellValue(makeParam, "E1", "Variant Price")

	for i, prod := range responseObjects {
		handleCell := "A" + strconv.Itoa(i+2)
		yearCell := "B" + strconv.Itoa(i+2)
		optionNameCell := "C" + strconv.Itoa(i+2)
		optionValueCell := "D" + strconv.Itoa(i+2)
		variantPriceCell := "E" + strconv.Itoa(i+2)
		mf.SetCellValue(makeParam, handleCell, prod.Handle)
		mf.SetCellValue(makeParam, yearCell, prod.Year)
		mf.SetCellValue(makeParam, optionNameCell, prod.OptionName)
		mf.SetCellValue(makeParam, optionValueCell, prod.OptionValue)
		mf.SetCellValue(makeParam, variantPriceCell, prod.Price)
	}
	mf.SetActiveSheet(mi)
	err := mf.SaveAs(makeParamStrFile)
	if err != nil {
		log.Println("Could not create xlsx file: ", err)
		return false
	} else {
		return true
	}

	// CSV - will be applied in future iterations as a flag.
	// f, err := os.Create(makeParamStrFile)

	// if err != nil {
	// 	log.Fatalln("failed to open file", err)
	// }

	// defer f.Close()

	// // Create the writer with the file
	// writer := csv.NewWriter(f)
	// defer writer.Flush()

	// csvRows := make([][]string, len(responseObjects))

	// for i, prod := range responseObjects {
	// 	csvRows[i] = []string{
	// 		prod.Handle,
	// 		strconv.FormatInt(prod.Year, 10),
	// 		prod.OptionName,
	// 		prod.OptionValue,
	// 		strconv.FormatFloat(prod.Price, 'f', 2, 32),
	// 	}
	// }

	// headerRow := []string{
	// 	"Handle",
	// 	"year",
	// 	"Option1 Name",
	// 	"Option1 Value",
	// 	"Variant Price",
	// }

	// tempRow := make([][]string, 1)

	// tempRow[0] = headerRow

	// csvRows = append(tempRow, csvRows...)
	// err = writer.WriteAll(csvRows) // calls Flush internally

	// if err != nil {
	// 	log.Fatal("Cannot write to file...", err)
	// 	return false
	// } else {
	// 	return true
	// }
}

//	product.Handle = "honda-accord-2000-touch-up-kit"
//	product.OptionName = "honda accord Paint Colors"
//	product.OptionValue = "Dark Emerald Pearl : G-87P"
//	product.Price = 39.21
//	product.Year = 2000
