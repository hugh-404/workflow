package states

import "context"

type SuccessState struct {
	StateCommon
}

func (s *SuccessState) Run(ctx context.Context) {
}