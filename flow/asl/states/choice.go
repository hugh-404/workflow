package states

type ChoiceState struct {
	StateCommon
	Choices []ChoiceTransition  `json:"Choices"`
}

type ChoiceTransition struct {
	Variable        string  `json:"Variable,omitempty"`
	StringEquals    string  `json:"StringEquals,omitempty"`
	NumericEquals   *int    `json:"NumericEquals,omitempty"`
	NumericGreaterThan *int  `json:"NumericGreaterThan,omitempty"`
	NumericLessThan *int     `json:"NumericLessThan,omitempty"`
}