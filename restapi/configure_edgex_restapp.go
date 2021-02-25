// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/momper14/edgex-restapp/util"
	"github.com/sirupsen/logrus"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/urfave/negroni"

	"github.com/momper14/edgex-restapp/restapi/operations"
)

//go:generate swagger generate server --target ../ --name edgex-restapp --spec ../openapi/iot-gateway-swagger-2.yaml --principal models.User --flag-strategy pflag

func configureFlags(api *operations.EdgexRestappAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.EdgexRestappAPI) http.Handler {
	// configure the api here
	api.ServeError = util.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	api.Logger = logrus.Infof

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the Authorization header is set with the Basic scheme
	api.BasicAuthAuth = authenticate

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	api.APIAuthorizer = authorizerFunc()

	configureDevices(api)
	configurePolicies(api)
	configureRoles(api)
	configureUsers(api)

	api.GetV1PingHandler = operations.GetV1PingHandlerFunc(
		func(params operations.GetV1PingParams) middleware.Responder {
			return operations.NewGetV1PingOK().WithPayload("pong")
		},
	)

	api.PreServerShutdown = func() {}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	logger := util.NewLogger()
	logger.ALogger = logrus.StandardLogger()
	logger.SetFormat(" {{.Status}} | {{.Duration}} | {{.Hostname}} | {{.Method}} {{.Path}}")

	recovery := negroni.NewRecovery()
	recovery.Logger = logrus.StandardLogger()

	n := negroni.New(logger, recovery)
	n.UseHandler(handler)
	return n
}
