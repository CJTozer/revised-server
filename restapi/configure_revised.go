package restapi

import (
	"crypto/tls"
	"net/http"
	"github.com/rs/cors"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"revised-server/models"

	"revised-server/restapi/operations"
	"revised-server/restapi/operations/books"
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

	api.BooksGetBooksHandler = books.GetBooksHandlerFunc(func(params books.GetBooksParams) middleware.Responder {
		// Dummy response for now
		booksList := []*models.Book{
			&models.Book{
				Author: "Dava Sobel",
				BookID: 1,
				Title:  "Longitude",
			},
			&models.Book{
				Author: "Robert K. Massie",
				BookID: 2,
				Title:  "Peter The Great",
			},
			&models.Book{
				Author: "Don Oberdorfer",
				BookID: 3,
				Title:  "The Two Koreas",
			},
		}

		return books.NewGetBooksOK().WithPayload(booksList)
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
