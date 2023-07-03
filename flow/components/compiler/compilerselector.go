package compiler

import (
	"encoding/json"
	"errors"

	"github.com/workflow/flow/asl/consts"
	"github.com/workflow/flow/asl/states"
)

type ICompiler interface {
	Compile(string, states.IState) (states.IState, error)
}

var compilerMap map[string]ICompiler
func init() {
	compilerMap = make(map[string]ICompiler)
	compilerMap["."] = &DefaultCompiler{}
}

type DefaultCompiler struct {
}

func (c *DefaultCompiler) Compile(asl string, state states.IState) (states.IState, error) {
	err := json.Unmarshal([]byte(asl), &state)
	if err != nil {
		return nil, err
	}
	return state, nil
}

func compile(stateMap map[string]interface{}) (states.IState, error) {
	stateBytes, _ := json.Marshal(stateMap)
	stateProto, err := getStateProto(stateMap["Type"].(string))
	if err != nil {
		return nil, err
	}
	return getCompiler(stateMap["Type"].(string)).Compile(string(stateBytes), stateProto)
}

func getCompiler(stateType string) ICompiler {
	if c, ok := compilerMap[stateType]; ok {
		return c
	}
	return compilerMap["."]
}

func getStateProto(stateType string) (states.IState, error) {
	var state states.IState
	switch stateType {
	case consts.State_Result:
		state = &states.ResultState{}
	case consts.State_Success:
		state = &states.SuccessState{}
	case consts.State_Task:
		state = &states.TaskState{}
	case consts.State_Choice:
		state = &states.ChoiceState{}

	}
	if state == nil {
		return nil, errors.New("invalid state type")
	}
	return state, nil
}