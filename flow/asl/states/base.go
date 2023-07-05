package states

import (
	"context"

	"github.com/workflow/flow/asl/execution"
)

type IState interface {
	Run(context.Context, *execution.ExecutionContext) (string, error)
	SetName(string)
	SetExecutionCtx(*execution.ExecutionContext)
	GetName() string
	GetExecutionCtx() *execution.ExecutionContext
}

type StateCommon struct {
	Name string
	Type         string
	Comment      string
	Param        map[string]interface{}
	Next         string
	InputPath    string
	OutputPath   string
	End bool
	ExecutionCtx *execution.ExecutionContext
}

func (c *StateCommon) SetName(name string) {
	c.Name = name
}
func (c *StateCommon) SetExecutionCtx(ctx *execution.ExecutionContext) {
	c.ExecutionCtx = ctx
}

func (c *StateCommon) GetName() string {
	return c.Name
}
func (c *StateCommon) GetExecutionCtx() *execution.ExecutionContext {
	return c.ExecutionCtx
}
