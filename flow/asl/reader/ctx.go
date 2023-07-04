package reader

import (
	"strings"

	"github.com/oliveagle/jsonpath"
	"github.com/workflow/flow/asl/execution"
)

type ReaderCtx string

func (c *ReaderCtx) Read(executionCtx *execution.ExecutionContext) (interface{}, error) {
	sc := string(*c)
	if strings.HasPrefix(sc, "$") && !strings.HasPrefix(sc, "$$") {
		// use local store
		jsonpath.JsonPathLookup(nil, sc)
	}
	
	return nil, nil
}