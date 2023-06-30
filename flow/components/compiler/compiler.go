package compiler

import (
	"encoding/json"
	"errors"

	"github.com/workflow/flow/asl/states"
)

type Compiler struct {

}

func (c *Compiler) Compile(asl string) (sm *states.StateMachine, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = errors.New("compile panic")
		}
	}()

	return unsafeCompile(asl)
}

func unsafeCompile(asl string) (*states.StateMachine, error) {
	tempMachine := struct {
		states.Base
		States map[string]map[string]interface{}
	}{}
	err := json.Unmarshal([]byte(asl), &tempMachine)
	if err != nil {
		return nil, err
	}
	sm := &states.StateMachine{}
	sm.Comment = tempMachine.Comment
	sm.StartAt = tempMachine.StartAt
	sm.Version = tempMachine.Version
	sm.TimeoutSeconds = tempMachine.TimeoutSeconds

	stateMap := make(map[string]states.IState)
	for name, stateObj := range tempMachine.States {
		stateBytes, _ := json.Marshal(stateObj)
		realState, err := UseStateCompiler(stateObj["Type"].(string), string(stateBytes))
		if err != nil {
			return nil, err
		}
		stateMap[name] = realState
	}

	sm.States = stateMap
	return sm, nil
}