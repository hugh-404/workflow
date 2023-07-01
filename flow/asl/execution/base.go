package execution

type ExecutionContext struct {
	GlobalResult map[string]interface{}
}

func (c *ExecutionContext)Init() {
	c.GlobalResult = map[string]interface{}{}
}
