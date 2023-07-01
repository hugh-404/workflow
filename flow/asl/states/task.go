package states

import (
	"context"
	"encoding/json"

	"github.com/workflow/flow/asl/execution"
)


type TaskState struct {
	StateCommon
	Resource string
}

func (s *TaskState)Run(ctx context.Context, exeCtx *execution.ExecutionContext) (string, error) {
	return "", nil
}

func (s *TaskState)Compile(asl string) (IState, error) {
	var state *TaskState
	err := json.Unmarshal([]byte(asl),&state)
	if err != nil {
		return nil, err
	}
	return state, nil
}

func CompileTask(asl string) (IState, error) {
	var state *TaskState
	err := json.Unmarshal([]byte(asl),&state)
	if err != nil {
		return nil, err
	}
	return state, nil
}