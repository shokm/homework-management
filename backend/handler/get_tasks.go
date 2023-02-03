package handler

import (
	"backend/gen/restapi/operations/task_api"

	"github.com/go-openapi/runtime/middleware"
)

func GetTasks(params task_api.GetTasksParams, principal interface{}) middleware.Responder {

	return middleware.NotImplemented("operation task_api.GetTasks has not yet been implemented")
}
