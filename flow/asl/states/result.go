package states

import (
	"context"
	"fmt"

	"github.com/workflow/flow/asl/consts"
	"github.com/workflow/flow/asl/execution"
)

type ResultState struct {
	StateCommon
}

func (s *ResultState) Run(ctx context.Context, exeCtx *execution.ExecutionContext) (string, error) {
	fmt.Println("run result")
	exeCtx.GlobalStore[consts.GlobalStore_Result] = s.Param["Result"]
	return s.Next, nil
}
