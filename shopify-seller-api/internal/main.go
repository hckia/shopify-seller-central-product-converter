package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/hckia/shopify-seller-central-product-converter/shopify-seller-api/pkg/swagger/server/models"
	"github.com/hckia/shopify-seller-central-product-converter/shopify-seller-api/pkg/swagger/server/restapi"

	"github.com/hckia/shopify-seller-central-product-converter/shopify-seller-api/pkg/swagger/server/restapi/operations"
)

/*
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Let's get to work, %q!", html.EscapeString(r.URL.Path))
	})

	log.Println("Listening to localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
*/

func main() {
	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln((err))
	}

	api := operations.NewHelloAPIAPI(swaggerSpec)

	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			// error hnadle
			log.Fatalln(err)
		}
	}()

	server.Port = 8080

	// following control statements in pkg/swagger/server/resapi/configure_hello_api.go
	// api.GetProductMakeHandler line 41
	// api.GetSwatchMakeHandler line 46
	// api.CheckHealthHandler line 51

	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)
	api.GetProductMakeHandler = operations.GetProductMakeHandlerFunc(GetProductMake)
	api.GetSwatchMakeHandler = operations.GetSwatchMakeHandlerFunc(GetSwatchMake)

	if err := server.Serve(); err != nil {
		log.Fatalln((err))
	}

}

//Health route returns okay
func Health(operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK().WithPayload("OK")
}

type productPayload struct {
	handle      string
	optionName  string
	optionValue string
	price       float32
	year        int
}

// comment is to initiate a PR for zeal to review accessor methods.

func GetProductMake(make operations.GetProductMakeParams) middleware.Responder {
	var URL string = ("https://someip.com/products/" + make.Make)
	fmt.Println(URL)  // https://someip.com/products/honda
	fmt.Println(make) // {0x14000598200 honda}
	//response, err := http.Get(URL)
	var err string = "some value"
	if err == "nil" {
		fmt.Println("The make provided does not exist, or some other error has occurred.")
	} else if strings.ToLower(make.Make) == "honda" {
		fmt.Println("Honda found.")
	}

	productPayloads := models.MakeProducts{}

	product := models.ProductRow{}
	product.Handle = "Goat"
	product.OptionName = "blah"
	product.OptionValue = "blah"
	product.Price = 39.21
	product.Year = 2000

	productPayloads = append(productPayloads, &product)

	return operations.NewGetProductMakeOK().WithPayload(productPayloads)
}

func GetSwatchMake(make operations.GetSwatchMakeParams) middleware.Responder {
	var URL string = ("https://someip.com/swatch/" + make.Make)
	fmt.Println(URL)  // https://someip.com/swatch/honda
	fmt.Println(make) // {0x14000598200 honda}
	//response, err := http.Get(URL)
	var err string = "some value"
	if err == "nil" {
		fmt.Println("The make provided does not exist, or some other error has occurred.")
	} else if strings.ToLower(make.Make) == "Honda" {
		fmt.Println("Honda found.")
	}

	swatchPayload := models.MakeSwatches{}

	swatch := models.SwatchRow{
		ColorCode: "",
		ColorName: "",
		Handle:    "",
		HexCode:   "",
		Make:      "",
		Mmy:       "",
		Model:     "",
		Tricoat:   false,
		Year:      0,
	}

	swatchPayload = append(swatchPayload, &swatch)

	return operations.NewGetSwatchMakeOK().WithPayload(swatchPayload)

}
