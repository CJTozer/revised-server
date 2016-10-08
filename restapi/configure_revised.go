package restapi

import (
	"crypto/tls"
	"net/http"
	"github.com/rs/cors"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"revised-server/restapi/operations"
	"revised-server/restapi/operations/books"
	"revised-server/restapi/operations/resources"

	"revised-server/backend"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ..\swagger.yaml

func configureFlags(api *operations.RevisedAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.RevisedAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Get the full list of books
	api.BooksGetBooksHandler = books.GetBooksHandlerFunc(func(params books.GetBooksParams) middleware.Responder {
		// Dummy response for now
		return books.NewGetBooksOK().WithPayload(backend.DummyBooksList())
	})

	// Get the full list of resources
	api.ResourcesGetResourcesHandler = resources.GetResourcesHandlerFunc(func(params resources.GetResourcesParams) middleware.Responder {
		// Dummy response for now
		return resources.NewGetResourcesOK().WithPayload(backend.DummyResourcesList())
	})

	// Get the resource with a specific ID
	api.ResourcesGetResourcesIDHandler = resources.GetResourcesIDHandlerFunc(func(params resources.GetResourcesIDParams) middleware.Responder {
		// Dummy response for now
		resList := backend.DummyResourcesList()
		if params.ID - 1 <= int64(len(resList)) {
			return resources.NewGetResourcesIDOK().WithPayload(resList[params.ID])
		}
		return resources.GetResourcesIDDefault(404)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
// Allow CORS.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
    handleCORS := cors.Default().Handler
    return handleCORS(handler)
}
