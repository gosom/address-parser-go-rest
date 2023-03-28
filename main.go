package main

import (
	"context"
	"embed"
	"fmt"
	"os"

	"github.com/gosom/kit/logging"
	"github.com/gosom/kit/web"
	"github.com/joho/godotenv"

	"github.com/gosom/address-parser-go-rest/addressparser/libpostal"
	"github.com/gosom/address-parser-go-rest/addressparser/ports"
)

//go:generate swag i -g main.go --pd

//go:embed docs/swagger.json
var specFs embed.FS

// @title Address Parser API
// @version 1.0.0
// @description This is the API for the address parser service

// @contact.name Giorgos Komninos
// @contact.url http://blog.gkomninos.com

// @host localhost:8080
// @BasePath /
// @accept json
// @produce json
// @query.collection.format multi
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	docPath := os.Getenv("DOCS_PATH")
	if docPath == "" {
		docPath = "/docs"
	}
	routerCfg := web.RouterConfig{
		SwaggerUI: &web.SwaggerUIConfig{
			SpecName: "AddressParser API",
			SpecFile: "/docs/swagger.json",
			Path:     docPath,
			SpecFS:   specFs,
		},
	}
	router := web.NewRouter(routerCfg)

	log := logging.Get()

	parser := libpostal.NewLibPostalParser(log.With("component", "libpostal"))
	hn := ports.NewAddressParserHandler(log.With("component", "handler"), parser)
	hn.RegisterRouters(router)
	addr := os.Getenv("PARSER_HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	webCfg := web.ServerConfig{
		Host:   addr,
		Router: router,
	}
	log.Info(fmt.Sprintf("Starting server at %s", addr))
	webSvc := web.NewHttpServer(webCfg)

	return webSvc.ListenAndServe(ctx)
}
