package sm

import (
	"context"

	"github.com/pkg/errors"
	"github.com/workflow/flow/asl/consts"
	"github.com/workflow/flow/asl/execution"
	"github.com/workflow/flow/asl/states"
)

type StateMachine struct {
	Base
	States map[string]states.IState
	ExecutionCtx *execution.ExecutionContext
}

type Base struct {
	Comment        string
	StartAt        string
	Version        string
	TimeoutSeconds int64
}

func (sm *StateMachine) SetExecutionCtx(ctx *execution.ExecutionContext) {
	sm.ExecutionCtx = ctx
	for _, state := range sm.States {
		state.SetExecutionCtx(ctx)
	}
}

func (sm *StateMachine) Run(ctx context.Context) error {
	stateMap := sm.States
	nextState := sm.StartAt
	var err error
	sm.ExecutionCtx.InputData[nextState] = sm.ExecutionCtx.GlobalStore[consts.GlobalStore_Input]
	for stateMap[nextState] != nil {
		curState := nextState
		nextState, err = sm.runState(ctx, stateMap[nextState])
		sm.ExecutionCtx.InputData[nextState] = sm.ExecutionCtx.OutputData[curState]
		if err != nil {
			return errors.Wrapf(err, "Run runState error: %v", curState)
		}
	}
	return nil
}

func (sm *StateMachine) runState(ctx context.Context, state states.IState) (string, error) {
	nextState, err := state.Run(ctx, sm.ExecutionCtx)
	if err != nil {
		return "", errors.Wrap(err, "runState Run error")
	}
	return nextState, nil
}