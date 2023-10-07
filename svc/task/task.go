package task

import "github.com/haileemiu/manage-life/svc/task/handler"

func NewHandler() *handler.Task {
	return handler.New()
}
