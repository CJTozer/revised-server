package main

import (
	"log"
	"os"
	"strconv"

	loads "github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"revised-server/restapi"
	"revised-server/restapi/operations"
)

// For swagger generation, it always overrides restapi/configure_revised.go, which we don't want
// to happen.  After much ado, simplest workaround is just to copy the original back afterwards.
//go:generate cp restapi/configure_revised.go restapi/configure_revised.go.pregenerate
//go:generate swagger generate server --spec swagger.yaml --exclude-main
//go:generate mv restapi/configure_revised.go.pregenerate restapi/configure_revised.go

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewRevisedAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = `revised API`
	parser.LongDescription = `Data API for revised-web.herokuapp.com`

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	// Serve the API as directed by the environment (and don't specify the host)
	port := os.Getenv("PORT")
	// Default port to 3000
	if port == "" {
		port = "3000"
	}
	server.Port, err = strconv.Atoi(port)
	if err != nil {
        log.Fatalln(err)
    }
	server.Host = ""
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
