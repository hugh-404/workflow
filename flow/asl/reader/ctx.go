package reader

import (
	"strings"

	"github.com/oliveagle/jsonpath"
	"github.com/workflow/flow/asl/states"
)

type ReaderCtx string

func (c *ReaderCtx) Read(state states.IState) (interface{}, error) {
	stateName := state.GetName()
	globalInput := state.GetExecutionCtx().GlobalStore["Input"]
	localInput := state.GetExecutionCtx().InputData[stateName]
	sc := string(*c)
	if strings.HasPrefix(sc, "$") {
		// is jsonpath
		if strings.HasPrefix(sc, "$$") {
			// use global context
			return jsonpath.JsonPathLookup(globalInput, sc[1:])
		} else {
			// use local context
			return jsonpath.JsonPathLookup(localInput, sc)
		}
	}
	
	return string(*c), nil
}