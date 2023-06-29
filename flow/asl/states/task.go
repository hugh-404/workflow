package states

import (
	"context"
	"encoding/json"
)


type TaskState struct {
	StateCommon
	Resource string
}

func (s *TaskState)Run(ctx context.Context) {

}

func (s *TaskState)Compile(asl string) (IState, error) {
	var state *TaskState
	err := json.Unmarshal([]byte(asl),&state)
	if err != nil {
		return nil, err
	}
	return state, nil
}