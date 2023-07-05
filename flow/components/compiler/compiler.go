package compiler

import (
	"encoding/json"
	"errors"

	"github.com/workflow/flow/asl/states"
	"github.com/workflow/flow/components/sm"
)

type Compiler struct {

}

func (c *Compiler) Compile(asl string) (sm *sm.StateMachine, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = errors.New("compile panic")
		}
	}()

	return unsafeCompile(asl)
}

func unsafeCompile(asl string) (*sm.StateMachine, error) {
	tempMachine := struct {
		sm.Base
		States map[string]map[string]interface{}
	}{}
	err := json.Unmarshal([]byte(asl), &tempMachine)
	if err != nil {
		return nil, err
	}
	sm := &sm.StateMachine{}
	sm.Comment = tempMachine.Comment
	sm.StartAt = tempMachine.StartAt
	sm.Version = tempMachine.Version
	sm.TimeoutSeconds = tempMachine.TimeoutSeconds

	stateMap := make(map[string]states.IState)
	for name, stateObj := range tempMachine.States {
		realState, err := compile(stateObj)
		if err != nil {
			return nil, err
		}
		realState.SetName(name)
		stateMap[name] = realState
	}

	sm.States = stateMap
	return sm, nil
}