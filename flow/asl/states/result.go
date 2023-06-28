package states

import "context"

type ResultState struct {
	StateCommon
}

type ResultStrategy string

const (
	Default = "Default"
	UseFix  = "UseFix"
	Merge   = "Merge"
)

func (s *ResultState) Run(ctx context.Context) {
	s.ExecutionCtx.GlobalResult["result"] = s.Param["Result"]
}
