package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
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

func GetProductMake(make operations.GetProductMakeParams) middleware.Responder {
	var URL string = ("https://someip.com/products?make=" + make.Make)

	response, err := http.Get(URL)

	if err != nil {
		fmt.Println("The make provided does not exist, or some other error has occurred.")
	}

	return operations.NewGetProductMakeOK().WithPayload(response.Body)
}

func GetSwatchMake(make operations.GetSwatchMakeParams) middleware.Responder {
	var URL string = ("https://someip.com/swatches?make=" + make.Make)

	response, err := http.Get(URL)

	if err != nil {
		fmt.Println("The make provided does not exist, or some other error has occurred.")
	}

	return operations.NewGetSwatchMakeOK().WithPayload(response.Body)
}
