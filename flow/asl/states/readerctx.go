package states

import (
	"encoding/json"
	"strings"

	"github.com/oliveagle/jsonpath"
	"github.com/workflow/flow/asl/consts"
)

type ReaderCtx string

func (c ReaderCtx) Read(state IState) (interface{}, error) {
	sc := string(c)
	stateName := state.GetName()
	globalInput := state.GetExecutionCtx().GlobalStore[consts.GlobalStore_Input]
	localInput := state.GetExecutionCtx().InputData[stateName]
	var globalInputMap, localInputMap interface{}
	globalJson, _ := json.Marshal(globalInput)
	UnmarshalNumber(string(globalJson), &globalInputMap)
	localJson, _ := json.Marshal(localInput)
	UnmarshalNumber(string(localJson), &localInputMap)
	if strings.HasPrefix(sc, "$") {
		// is jsonpath
		if strings.HasPrefix(sc, "$$") {
			// use global context
			return jsonpath.JsonPathLookup(globalInputMap, sc[1:])
		} else {
			// use local context
			return jsonpath.JsonPathLookup(localInputMap, sc)
		}
	}
	
	return c, nil
}

func UnmarshalNumber(jsonStr string, result interface{}) error {
	decoder := json.NewDecoder((strings.NewReader(jsonStr)) )
 	decoder.UseNumber() 
 	return decoder.Decode(&result) 
}
