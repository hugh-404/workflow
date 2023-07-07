package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func TestFileRead(t *testing.T) {
	bs, err := ioutil.ReadFile("/var/asl/aslmap.json")
	if err != nil {
		fmt.Println(err)
	}
	var aslMap map[string]interface{}
	err = json.Unmarshal(bs, &aslMap)
	if err != nil {
		fmt.Println(err)
	}
	v, err := json.Marshal(aslMap["key1"])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(v))
}