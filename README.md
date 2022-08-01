# shopify-seller-central-product-converter

This repository is a prototype to convert shopify products from timefortouchup.com into flat files that can be utilized to create content and products for Amazon seller central. While the production of the content creation and final excel file is contained in a private repository, this golang application will produce csv and excel files that will contain the necessary data.

There are two interfaces to this prototype.

## The REST API (shopify-seller-api directory)

**Information on the methods can be found within the shopify-seller-api/doc/index.html file.**

- The initial goal is for the API to only provide read access to the data. Users will not be able to to Create, Update, or Delete.
- Since this will most likely be localized, authentication may not be considered. If it is hosted on a server all IPs will be blocked except for an allowed list.

## The CLI (shopify-seller-cli directory)

- The CLI may start off utilizing conditional logic as a rudementary interface. If time Permits, Cobra will be utilized. 
- If Cobra is utilized, Viper will likely be incorporated as well.

## Other things to note

I'm drawing inspiration from an CFR creator utilized by peers, along with [the following tutorial](https://dev.to/aurelievache/learning-go-by-examples-part-2-create-an-http-rest-api-server-in-go-1cdm), that I have been utilizing to as a boilerplate for my taskfile.yml, and swagger.yml
