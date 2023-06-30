package states

import (
	"context"

	"github.com/workflow/flow/asl/consts"
)

type ResultState struct {
	StateCommon
}

func (s *ResultState) Run(ctx context.Context) {
	s.ExecutionCtx.GlobalResult[consts.GlobalResult_Result] = s.Param["Result"]
}
