package compiler

import (
	"encoding/json"
	"errors"

	"github.com/workflow/flow/asl/states"
)

func UseStateCompiler(stateType, asl string) (states.IState, error) {
	var state states.IState
	switch stateType {
	case states.State_Task:
		state = &states.TaskState{}
	default:
		return nil, errors.New("no such state")
	}
	err := json.Unmarshal([]byte(asl), &state)
	if err != nil {
		return nil, err
	}
	return state, nil
}