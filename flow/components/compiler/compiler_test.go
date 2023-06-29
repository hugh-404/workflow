package compiler

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCompile(t *testing.T) {
	asl := `
	{
		"Comment": "A simple minimal example of the States language",
		"StartAt": "Hello World",
		"States": {
		  "Hello World": {
			"Type": "Task",
			"Resource": "arn:aws:lambda:us-east-1:123456789012:function:HelloWorld",
			"End": true
		  }
		}
	  }`
	sm, err := (&Compiler{}).Compile(asl)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, sm, nil)
	assert.Equal(t, len(sm.States), 1)
}
