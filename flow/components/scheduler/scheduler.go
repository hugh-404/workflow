package scheduler

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/workflow/flow/asl/consts"
	"github.com/workflow/flow/asl/execution"
	"github.com/workflow/flow/components/compiler"
)

type Scheduler struct {}

func (s *Scheduler) StartExecution(ctx context.Context, asl string, params map[string]interface{}) (*execution.ExecutionContext, error) {
	fsm, err := (&compiler.Compiler{}).Compile(asl)
	if err != nil {
		fmt.Println("compile error")
		return nil, errors.Wrap(err, "StartExecution Compile error")
	}
	fsm.SetExecutionCtx(&execution.ExecutionContext{})
	fsm.ExecutionCtx.Init()
	fsm.ExecutionCtx.GlobalStore[consts.GlobalStore_Input] = params
	err = fsm.Run(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "StartExecution Run error")
	}
	return  fsm.ExecutionCtx, nil
}

