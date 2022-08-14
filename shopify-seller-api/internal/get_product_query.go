package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/hckia/shopify-seller-central-product-converter/shopify-seller-api/pkg/swagger/server/models"
)

func getProductQuery(make string) []*models.ProductRow {

	db, err := sql.Open("mysql", os.Getenv("T4TDBCONSTRING"))

	if err != nil {
		log.Fatalln("ERROR ", err.Error())
	}

	defer db.Close()

	log.Println("Successfully connected to MySQL Database")
	//	product.Handle = "honda-accord-2000-touch-up-kit"
	//	product.OptionName = "honda accord Paint Colors"
	//	product.OptionValue = "Dark Emerald Pearl : G-87P"
	//	product.Price = 39.21
	//	product.Year = 2000

	qStr := "SELECT `Handle`, `Option1 Name`,  `Option1 Value`, `Variant Price`, CAST(`Year` AS SIGNED) FROM product_makes_combined_data WHERE `Metafield: product.make [single_line_text_field]` = '" + make + "';"
	found, errr := db.Query(qStr)

	if errr != nil {
		log.Fatalln("ERROR: ", errr.Error())
	}

	counter := 1
	var productPayloads []*models.ProductRow

	for found.Next() {
		makeFound := models.ProductRow{}

		err = found.Scan(&makeFound.Handle, &makeFound.OptionName, &makeFound.OptionValue, &makeFound.Price, &makeFound.Year)
		if err != nil {
			log.Fatalln("ERR ", err.Error())
		}

		productPayloads = append(productPayloads, &makeFound)
		counter += 1
	}
	log.Printf("Number of rows found: %d", counter)

	return productPayloads
}
