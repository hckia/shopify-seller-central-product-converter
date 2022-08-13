package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/hckia/shopify-seller-central-product-converter/shopify-seller-api/pkg/swagger/server/models"
)

func getSwatchQuery(make string) []*models.SwatchRow {

	db, err := sql.Open("mysql", os.Getenv("T4TDBCONSTRING"))

	if err != nil {
		log.Fatalln("ERROR ", err.Error())
	}

	defer db.Close()

	log.Println("Successfully connected to MySQL Database")

	qStr := "SELECT `Color code`, `Color name`, `Handle`, `Variant Metafield: variants.color [color]`, `Make`, `MMY`, `Model`, IF(`Tricoat`=1,'true','false'), CAST(`Year` AS SIGNED) FROM swatch_hex_data WHERE Make = '" + make + "';"
	//fmt.Println(qStr)
	found, errr := db.Query(qStr)

	if errr != nil {
		log.Fatalln("ERROR: ", errr.Error())
	}

	counter := 1
	var swatchPayload []*models.SwatchRow

	for found.Next() {
		makeFound := models.SwatchRow{}
		//
		err = found.Scan(&makeFound.ColorCode, &makeFound.ColorName, &makeFound.Handle, &makeFound.HexCode, &makeFound.Make, &makeFound.Mmy, &makeFound.Model, &makeFound.Tricoat, &makeFound.Year)
		if err != nil {
			log.Fatalln("ERR ", err.Error())
		}

		swatchPayload = append(swatchPayload, &makeFound)
		counter += 1
	}
	log.Printf("Number of rows found: %d", counter)

	return swatchPayload
}
