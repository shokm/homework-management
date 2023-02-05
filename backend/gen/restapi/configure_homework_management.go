// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/rs/cors"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"backend/gen/restapi/operations"
	"backend/gen/restapi/operations/auth_api"
	"backend/gen/restapi/operations/task_api"
	"backend/handler"
	"backend/handler/auth_jwt"
)

//go:generate swagger generate server --target ../../gen --name HomeworkManagement --spec ../../swagger/swagger.yaml --principal interface{}

func configureFlags(api *operations.HomeworkManagementAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HomeworkManagementAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	/*
	if api.BearerAuthAuth == nil {
		api.BearerAuthAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (bearerAuth) Authorization from header param [Authorization] has not yet been implemented")
		}
	}
	*/
	api.BearerAuthAuth = auth_jwt.ValidateTokenHandler

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	/*
	if api.AuthAPIPostAuthLoginHandler == nil {
		api.AuthAPIPostAuthLoginHandler = auth_api.PostAuthLoginHandlerFunc(func(params auth_api.PostAuthLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth_api.PostAuthLogin has not yet been implemented")
		})
	}
	*/
	api.AuthAPIPostAuthLoginHandler = auth_api.PostAuthLoginHandlerFunc(handler.PostAuthLogin)

	/*
	if api.AuthAPIPostAuthRegisterHandler == nil {
		api.AuthAPIPostAuthRegisterHandler = auth_api.PostAuthRegisterHandlerFunc(func(params auth_api.PostAuthRegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation auth_api.PostAuthRegister has not yet been implemented")
		})
	}
	*/
	api.AuthAPIPostAuthRegisterHandler = auth_api.PostAuthRegisterHandlerFunc(handler.PostAuthRegister)

	/*
	if api.AuthAPIGetAuthUserHandler == nil {
		api.AuthAPIGetAuthUserHandler = auth_api.GetAuthUserHandlerFunc(func(params auth_api.GetAuthUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation auth_api.GetAuthUser has not yet been implemented")
		})
	}
	*/
	api.AuthAPIGetAuthUserHandler = auth_api.GetAuthUserHandlerFunc(handler.GetAuthUser)

	/*
	if api.TaskAPIGetSubjectBySubjectIDHandler == nil {
		api.TaskAPIGetSubjectBySubjectIDHandler = task_api.GetSubjectBySubjectIDHandlerFunc(func(params task_api.GetSubjectBySubjectIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation task_api.GetSubjectBySubjectID has not yet been implemented")
		})
	}
	*/
	api.TaskAPIGetSubjectBySubjectIDHandler = task_api.GetSubjectBySubjectIDHandlerFunc(handler.GetSubjectBySubjectID)

	/*
	if api.TaskAPIGetSubjectsHandler == nil {
		api.TaskAPIGetSubjectsHandler = task_api.GetSubjectsHandlerFunc(func(params task_api.GetSubjectsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation task_api.GetSubjects has not yet been implemented")
		})
	}
	*/
	api.TaskAPIGetSubjectsHandler = task_api.GetSubjectsHandlerFunc(handler.GetSubjects)

	/*
	if api.TaskAPIGetTaskByTaskIDHandler == nil {
		api.TaskAPIGetTaskByTaskIDHandler = task_api.GetTaskByTaskIDHandlerFunc(func(params task_api.GetTaskByTaskIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation task_api.GetTaskByTaskID has not yet been implemented")
		})
	}
	*/
	api.TaskAPIGetTaskByTaskIDHandler = task_api.GetTaskByTaskIDHandlerFunc(handler.GetTaskByTaskID)

	/*
	if api.TaskAPIGetTasksHandler == nil {
		api.TaskAPIGetTasksHandler = task_api.GetTasksHandlerFunc(func(params task_api.GetTasksParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation task_api.GetTasks has not yet been implemented")
		})
	}
	*/
	api.TaskAPIGetTasksHandler = task_api.GetTasksHandlerFunc(handler.GetTasks)

	/*
	if api.TaskAPIGetTasksBySubjectsHandler == nil {
		api.TaskAPIGetTasksBySubjectsHandler = task_api.GetTasksBySubjectsHandlerFunc(func(params task_api.GetTasksBySubjectsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation task_api.GetTasksBySubjects has not yet been implemented")
		})
	}
	*/
	api.TaskAPIGetTasksBySubjectsHandler = task_api.GetTasksBySubjectsHandlerFunc(handler.GetTasksBySubjects)

	/*
	if api.TaskAPIPostSubjectBySubjectIDHandler == nil {
		api.TaskAPIPostSubjectBySubjectIDHandler = task_api.PostSubjectBySubjectIDHandlerFunc(func(params task_api.PostSubjectBySubjectIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation task_api.PostSubjectBySubjectID has not yet been implemented")
		})
	}
	*/
	api.TaskAPIPostSubjectBySubjectIDHandler = task_api.PostSubjectBySubjectIDHandlerFunc(handler.PostSubjectBySubjectID)

	/*
	if api.TaskAPIPostTaskByTaskIDHandler == nil {
		api.TaskAPIPostTaskByTaskIDHandler = task_api.PostTaskByTaskIDHandlerFunc(func(params task_api.PostTaskByTaskIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation task_api.PostTaskByTaskID has not yet been implemented")
		})
	}
	*/
	api.TaskAPIPostTaskByTaskIDHandler = task_api.PostTaskByTaskIDHandlerFunc(handler.PostTaskByTaskID)

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
	/*
	return handler
	*/

	// TODO: CORSの設定をする
	/*
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://foo.com"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	/**/

	handleCORS := cors.Default().Handler

    return handleCORS(handler)
}
