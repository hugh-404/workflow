package execution

type ExecutionContext struct {
	StateRelation
	GlobalStore map[string]interface{}
	InputData map[string]interface{}
	OutputData map[string]interface{}
}

type StateRelation struct {
	Parent string
	Precending string
}

func (c *ExecutionContext)Init() {
	c.GlobalStore = map[string]interface{}{}
	c.InputData = map[string]interface{}{}
	c.OutputData = map[string]interface{}{}
}
