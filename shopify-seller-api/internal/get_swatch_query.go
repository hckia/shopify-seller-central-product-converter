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
	found, errr := db.Query(qStr)

	if errr != nil {
		log.Fatalln("ERROR: ", errr.Error())
	}

	counter := 1
	var swatchPayload []*models.SwatchRow

	for found.Next() {
		makeFound := models.SwatchRow{}
		// swatch := models.SwatchRow{
		// 	ColorCode: "G-87P",
		// 	ColorName: "Dark Emerald Pearl",
		// 	Handle:    "honda-accord-2000-touch-up-kit",
		// 	HexCode:   "#0c5c1c",
		// 	Make:      "Honda",
		// 	Mmy:       "2000 Honda Accord",
		// 	Model:     "Accord",
		// 	Tricoat:   false,
		// 	Year:      2000,
		// }
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
