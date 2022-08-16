# shopify-seller-central-product-converter

This monorepo is a prototype to convert shopify products from timefortouchup.com into flat files that can be utilized to create content and products for Amazon seller central. While the production of the content creation and final excel file is contained in a private repository, this golang application will produce csv and excel files that will contain the necessary data.

There are two interfaces to this prototype.

## The REST API (shopify-seller-api directory)

**Information on the methods can be found within the shopify-seller-api/doc/index.html file.**

- The the API only provides read access to the data. Users will not be able to to Create, Update, or Delete for now.

## sscli (the shopify-seller-cli directory)

- To install sscli executable users will have to cd into sscli and run `go build -o bin/sscli main.go` 
- CLI was built utilizing Cobra with the following commands
    - fetchProducts 
        - Description: Fetches an xlsx file with the make of your choosing from a personally developed vehicle paint database. It is utilized in a python script (also developed by me, but private) for timefortouchup.com to upload products to Amazon Seller Central
        - args makeParam: make of the vehicle
        - args path: path to deploy the xlsx file (e.g. /Chosen/Path/Data/Product/Make)
    - fetchSwatches 
        - Description: Fetches a csv file with the make of your choosing. It is utilized in a python script (also developed by me, but private) for timefortouchup.com to upload color swatches to Amazon Seller Central
        - args makeParam: make of the vehicle
        - args path: path to deploy the csv file (e.g. /Chosen/Path/Data/Swatch/Make)
## Other things to note

I'm drawing inspiration from an CFR creator utilized by peers, along with [the following tutorial](https://dev.to/aurelievache/learning-go-by-examples-part-2-create-an-http-rest-api-server-in-go-1cdm), that I have been utilizing to as a boilerplate for my taskfile.yml, and swagger.yml
