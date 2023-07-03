package states

import (
	"context"

	"github.com/workflow/flow/asl/execution"
)

type ChoiceState struct {
	StateCommon
	Branches []Branch  `json:"Branches"`
}

type Branch struct {
	ChoiceComparator
	Next string `json:"Next,omitempty"`
}

type ChoiceComparator struct {
	Variable        *string  `json:"Variable,omitempty"`
	StringEquals    *string  `json:"StringEquals,omitempty"`
	NumericEquals   *int    `json:"NumericEquals,omitempty"`
	NumericGreaterThan *int  `json:"NumericGreaterThan,omitempty"`
	StringIncludes *string `json:"StringIncludes,omitempty"`
	StringEndWith *string `json:"StringEndWith,omitempty"`
	Boolean *bool `json:"Boolean,omitempty"`
	And []*ChoiceComparator `json:"And,omitempty"`
	Or []*ChoiceComparator `json:"Or,omitempty"`
	Not *ChoiceComparator `json:"Not,omitempty"`
}

func (s *ChoiceState) Run(ctx context.Context, exeCtx *execution.ExecutionContext) (string, error) {
	return "", nil
}