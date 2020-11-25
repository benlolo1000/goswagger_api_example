// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"

	_ "github.com/go-sql-driver/mysql"

	"medicarePartD/gen/restapi/operations"
	"medicarePartD/gen/restapi/operations/login"
	"medicarePartD/gen/restapi/operations/record"
	"medicarePartD/gen/restapi/operations/records"
	"medicarePartD/gen/restapi/operations/register"
)

//go:generate swagger generate server --target ../../medicareApiClone --name Medicare --spec ../swagger/swagger.yml

func configureFlags(api *operations.MedicareAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MedicareAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.RecordsGetRecordsListHandler = records.GetRecordsListHandlerFunc(func(params records.GetRecordsListParams, principal interface{}) middleware.Responder {
		return records.HandleRecordsList(params)
	})
	api.RecordGetSingleRecordHandler = record.GetSingleRecordHandlerFunc(func(params record.GetSingleRecordParams, principal interface{}) middleware.Responder {
		return record.HandleSingleRecord(params)
	})
	api.LoginLoginHandler = login.LoginHandlerFunc(func(params login.LoginParams) middleware.Responder {
		return login.HandleLogin(params)
	})
	api.RegisterRegisterHandler = register.RegisterHandlerFunc(func(params register.RegisterParams) middleware.Responder {
		return register.HandleRegister(params)
	})
	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {

	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST"},
		AllowCredentials: true,
	})
	return handleCORS.Handler(handler)
}
