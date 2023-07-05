package execution

type ExecutionContext struct {
	GlobalStore map[string]interface{}
	InputData map[string]interface{}
	OutputData map[string]interface{}
}

func (c *ExecutionContext)Init() {
	c.GlobalStore = map[string]interface{}{}
	c.InputData = map[string]interface{}{}
	c.OutputData = map[string]interface{}{}
}
