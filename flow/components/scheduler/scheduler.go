package scheduler

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/workflow/flow/asl/execution"
	"github.com/workflow/flow/components/compiler"
)

type Scheduler struct {}

func (s *Scheduler) StartExecution(ctx context.Context, asl string) (*execution.ExecutionContext, error) {
	fsm, err := (&compiler.Compiler{}).Compile(asl)
	if err != nil {
		fmt.Println("compile error")
		return nil, errors.Wrap(err, "StartExecution Compile error")
	}
	fsm.ExecutionCtx = &execution.ExecutionContext{}
	fsm.ExecutionCtx.Init()
	err = fsm.Run(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "StartExecution Run error")
	}
	return  fsm.ExecutionCtx, nil
}

