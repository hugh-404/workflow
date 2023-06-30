package states

import (
	"context"

	"github.com/workflow/flow/asl/execution"
)

type StateMachine struct {
	Base
	States map[string]IState
}

type Base struct {
	Comment        string
	StartAt        string
	Version        string
	TimeoutSeconds int64
}

type IState interface {
	Run(context.Context)
	Compile(string) (IState, error)
}

type StateCommon struct {
	Type         string
	Comment      string
	Param        map[string]interface{}
	Next         string
	InputPath    string
	OutputPath   string
	ExecutionCtx *execution.ExecutionContext
	End bool
}

const (
	State_Task = "Task"
	State_Result = "Result"
	State_Success = "Success"
	State_Choice = "Choice"
	State_Parallel = "Parallel"
)