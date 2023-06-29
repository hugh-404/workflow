package compiler

import (
	"errors"

	"github.com/workflow/flow/asl/states"
)

func UseStateCompiler(stateType string) (states.IState, error) {
	switch stateType {
	case states.State_Task:
		return &states.TaskState{}, nil
	}
	return nil, errors.New("no such state")
}