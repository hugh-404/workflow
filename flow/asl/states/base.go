package states

import (
	"context"

	"github.com/workflow/flow/asl/execution"
)

type IState interface {
	Run(context.Context, *execution.ExecutionContext) (string, error)
}

type StateCommon struct {
	Type         string
	Comment      string
	Param        map[string]interface{}
	Next         string
	InputPath    string
	OutputPath   string
	End bool
}
