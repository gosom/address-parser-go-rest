# Address Parser Go REST
[![Go Report Card](https://goreportcard.com/badge/github.com/gosom/address-parser-go-rest)](https://goreportcard.com/report/github.com/gosom/address-parser-go-rest)


Address Parser Go REST is a REST API that provides address parsing functionality using the libpostal library. 
The purpose of this API is to allow users to easily parse addresses into their individual components 
without the need for the libpostal library to be included as a dependency in their projects.

## Quickstart

```
docker run -p 8080 gosom/address-parser-go-rest:v1.0.0
```

This will take some time to load

then try a sample request

```
curl -X 'POST' \
  'http://localhost:8080/parse' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "address": "48 Leicester Square, London WC2H 7LU, United Kingdom",
   "title_case": true
}'
```

Response:

```
curl -X 'POST' \
  'http://localhost:8080/parse' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "address": "48 Leicester Square, London WC2H 7LU, United Kingdom",
   "title_case": true
}'
```

[swagger documentation](http://localhost:8080/docs/)


## Run without docker

To install and run Address Parser Go REST, you can use the following steps:

1. Make sure you have a recent version of Golang
2. [Install](https://github.com/openvenues/libpostal/issues#installation-maclinux) libpostal on your machine.
3. `go mod tidy`
4. `go run main.go`
   

Notes:
you can change the port the service or the path for swagger is listening to by setting the following environment variables:
```
PARSER_HTTP_ADDR=:8080
DOCS_PATH=/docs
```
you can also put these in `.env` file in the root of the project.

If you want to rebuild the swagger documentation make sure that you have
installed [swag](https://github.com/swaggo/swag)

to regenerate:
```
go generate
```

## Contributing

If you would like to contribute to Address Parser Go REST, please create a pull request with your changes. 
You can also report any issues or bugs you encounter by creating a new issue on the GitHub repository.

## License

Address Parser Go REST is licensed under the MIT License. See `LICENSE` for more information.

## Acknowledgments

We would like to acknowledge the contributors of the libpostal library and the Go bindings used in this project.


