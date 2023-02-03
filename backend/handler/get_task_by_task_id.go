package handler

import (
	"backend/gen/restapi/operations/task_api"

	"github.com/go-openapi/runtime/middleware"
)

func GetTaskByTaskID(params task_api.GetTaskByTaskIDParams, principal interface{}) middleware.Responder {

	return middleware.NotImplemented("operation task_api.GetTaskByTaskID has not yet been implemented")
}
