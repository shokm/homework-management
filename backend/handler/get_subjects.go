package handler

import (
	"backend/gen/restapi/operations/task_api"

	"github.com/go-openapi/runtime/middleware"
)

func GetSubjects(params task_api.GetSubjectsParams, principal interface{}) middleware.Responder {

	return middleware.NotImplemented("operation task_api.GetSubjects has not yet been implemented")
}
