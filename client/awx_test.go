package client

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestAwx(t *testing.T) {
	a, _ := (&AwxCient{}).Run(context.TODO(), map[string]interface{}{
		"Var1": 10,
		"Var2": "Success",
		"Var3": struct{
			InnerS string
		}{
			InnerS: "XXX",
		},
	})
	v, _ := json.Marshal(a)
	fmt.Println(string(v))
	time.Sleep(500*time.Millisecond)
}