package main

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hckia/shopify-seller-central-product-converter/shopify-seller-api/pkg/swagger/server/restapi"

	"github.com/hckia/shopify-seller-central-product-converter/shopify-seller-api/pkg/swagger/server/restapi/operations"
)

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
	// var URL string = ("/product/" + make.Make) // if we went to another endpoint // https://someip.com/product/honda
	// log.Println(make) // {0x14000598200 honda}
	//response, err := http.Get(URL)

	var productPayload = getProductQuery(make.Make)

	return operations.NewGetProductMakeOK().WithPayload(productPayload)
}

func GetSwatchMake(make operations.GetSwatchMakeParams) middleware.Responder {
	// var URL string = ("/swatch/" + make.Make) // if we went to another endpoint // https://someip.com/swatch/honda
	// log.Println(make) // {0x14000598200 honda}
	//response, err := http.Get(URL)

	var swatchPayload = getSwatchQuery(make.Make)

	return operations.NewGetSwatchMakeOK().WithPayload(swatchPayload)

}
