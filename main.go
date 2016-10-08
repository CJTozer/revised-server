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

//go:generate swagger generate server --spec swagger.yaml

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
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalln(err)
	}
	server.Port = port
	server.Host = ""
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
