package states

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/workflow/flow/asl/execution"
)

type ChoiceState struct {
	StateCommon
	Branches []Branch  `json:"Branches,omitempty"`
	Default string	`json:"Default,omitempty"`
}

type Branch struct {
	ChoiceComparator
	Next string `json:"Next,omitempty"`
}

type ChoiceComparator struct {
	Variable        *ReaderCtx  `json:"Variable,omitempty"`
	StringEquals    *StringEquals  `json:"StringEquals,omitempty"`
	NumericEquals   *NumericEquals    `json:"NumericEquals,omitempty"`
	NumericGreaterThan *NumericGreaterThan  `json:"NumericGreaterThan,omitempty"`
	StringContains *StringContains `json:"StringContains,omitempty"`
	StringEndWith *StringEndWith `json:"StringEndWith,omitempty"`
	Boolean *BooleanT `json:"Boolean,omitempty"`
	And []*ChoiceComparator `json:"And,omitempty"`
	Or []*ChoiceComparator `json:"Or,omitempty"`
	Not *ChoiceComparator `json:"Not,omitempty"`
}

type StringEquals string
func (c *StringEquals) Compare(operand interface{}) bool {
	if s, ok := operand.(string); !ok {
		panic("StringEquals invalid state")
	} else {
		return s == string(*c)
	}
}

type StringContains string
func (c *StringContains) Compare(operand interface{}) bool {
	if s, ok := operand.(string); !ok {
		panic("StringContains invalid state")
	} else {
		return strings.Contains(s, string(*c))
	}
}

type StringEndWith string
func (c *StringEndWith) Compare(operand interface{}) bool {
	if s, ok := operand.(string); !ok {
		panic("StringEndWith invalid state")
	} else {
		return strings.HasSuffix(s, string(*c))
	}
}

type NumericEquals int
func (c *NumericEquals) Compare(operand interface{}) bool {
	if s, ok := operand.(json.Number); !ok {
		panic("NumericEquals invalid state")
	} else {
		n, _ := s.Int64()
		return n == int64(*c)
	}
}

type NumericGreaterThan int
func (c *NumericGreaterThan) Compare(operand interface{}) bool {
	if s, ok := operand.(json.Number); !ok {
		panic("NumericGreaterThan invalid state")
	} else {
		n, _ := s.Int64()
		return n > int64(*c)
	}
}

type BooleanT bool
func (c *BooleanT) Compare(operand interface{}) bool {
	if s, ok := operand.(bool); !ok {
		panic("BooleanT invalid state")
	} else {
		return s == bool(*c)
	}
}

func (s *ChoiceState) Run(ctx context.Context, exeCtx *execution.ExecutionContext) (nextState string, err error) {
	fmt.Println("run in ChoiceState")
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("ChoiceState error")
		}
	}()
	for _, b := range s.Branches {
		if b.ChoiceComparator.Compare(ctx, s) {
			return b.Next, nil
		}
	}
	return s.Default, nil
}

func (c *ChoiceComparator)Compare(ctx context.Context, state IState) bool {
	if c.Variable != nil {
		// compare directly
		realVar, err := c.Variable.Read(state)
		if err != nil {
			realVar = string(*c.Variable)
		}
		if c.StringEquals != nil {
			return c.StringEquals.Compare(realVar)
		} else if c.StringContains != nil {
			return c.StringContains.Compare(realVar)
		} else if c.StringEndWith != nil {
			return c.StringEndWith.Compare(realVar)
		} else if c.NumericEquals != nil {
			return c.NumericEquals.Compare(realVar)
		} else if c.NumericGreaterThan != nil {
			return c.NumericGreaterThan.Compare(realVar)
		} else if c.Boolean != nil {
			return c.Boolean.Compare(realVar)
		}
	} else {
		// use and|not|or comparator
		if len(c.And) > 0 {
			for _, ic := range c.And {
				if !ic.Compare(ctx, state) {
					return false
				}
			}
			return true
		}
		if len(c.Or) > 0 {
			for _, ic := range c.Or {
				if ic.Compare(ctx, state) {
					return true
				}
			}
			return false
		}
		if c.Not != nil {
			return !c.Compare(ctx, state)
		}
	}
	panic("invalid choice state")
}