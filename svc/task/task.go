package task

import (
	"github.com/haileemiu/manage-life/ent"
	"github.com/haileemiu/manage-life/svc/task/handler"
)

func NewHandler(entClient *ent.Client) *handler.Task {
	return handler.New(entClient)
}
