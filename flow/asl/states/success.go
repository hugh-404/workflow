package states

import (
	"context"
	"fmt"

	"github.com/workflow/flow/asl/execution"
)

type SuccessState struct {
	StateCommon
}

func (s *SuccessState) Run(ctx context.Context, exeCtx *execution.ExecutionContext) (string, error) {
	fmt.Println("run success")
	return "", nil
}