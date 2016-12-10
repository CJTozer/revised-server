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

	// === GET REQUESTS ===

	// Get the full list of books
	api.BooksGetBooksHandler = books.GetBooksHandlerFunc(func(params books.GetBooksParams) middleware.Responder {
		// Dummy response for now
		rspPayload := books.GetBooksOKBody{backend.DummyBooksList}
		return books.NewGetBooksOK().WithPayload(rspPayload)
	})

	// Get the book with a specific ID
	api.BooksGetBooksIDHandler = books.GetBooksIDHandlerFunc(func(params books.GetBooksIDParams) middleware.Responder {
		// Dummy response for now
		booksList := backend.DummyBooksList
		if params.ID - 1 <= int64(len(booksList)) {
			rspPayload := books.GetBooksIDOKBody{booksList[params.ID - 1]}
			return books.NewGetBooksIDOK().WithPayload(rspPayload)
		}
		return books.NewGetBooksIDDefault(404)
	})

	// Get the full list of resources
	api.ResourcesGetResourcesHandler = resources.GetResourcesHandlerFunc(func(params resources.GetResourcesParams) middleware.Responder {
		// Dummy response for now
		rspPayload := resources.GetResourcesOKBody{backend.DummyResourcesList()}
		return resources.NewGetResourcesOK().WithPayload(rspPayload)
	})

	// Get the resource with a specific ID
	api.ResourcesGetResourcesIDHandler = resources.GetResourcesIDHandlerFunc(func(params resources.GetResourcesIDParams) middleware.Responder {
		// Dummy response for now
		resList := backend.DummyResourcesList()
		if params.ID - 1 <= int64(len(resList)) {
			rspPayload := resources.GetResourcesIDOKBody{resList[params.ID - 1]}
			return resources.NewGetResourcesIDOK().WithPayload(rspPayload)
		}
		return resources.NewGetResourcesIDDefault(404)
	})

	// === POST REQUESTS ===

	// Post a new book
	api.BooksPostBooksHandler = books.PostBooksHandlerFunc(func(params books.PostBooksParams) middleware.Responder {
		// Set the ID, save the new book and return it
		params.Book.Book.ID = backend.GetNextBookID()
		backend.DummyBooksList = append(backend.DummyBooksList, params.Book.Book)
		rspPayload := books.PostBooksOKBody{params.Book.Book}
		return books.NewPostBooksOK().WithPayload(rspPayload)
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
