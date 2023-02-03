package handler

import (
	"backend/gen/restapi/operations/task_api"

	"github.com/go-openapi/runtime/middleware"
)

func GetTasksBySubjects(params task_api.GetTasksBySubjectsParams, principal interface{}) middleware.Responder {

	return middleware.NotImplemented("operation task_api.GetTasksBySubjects has not yet been implemented")
}
