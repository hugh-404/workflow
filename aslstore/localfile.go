package aslstore

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/workflow/flow/asl/consts"
)

type LocalFile struct{}

func (f *LocalFile) Fetch(ctx context.Context, key string) (asl string) {
	defer func() {
		if asl == "" {
			asl = consts.SimpleResultAsl
		}
	}()
	bs, err := ioutil.ReadFile("/var/asl/aslmap.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var aslMap map[string]interface{}
	err = json.Unmarshal(bs, &aslMap)
	if err != nil {
		fmt.Println(err)
	}
	v, err := json.Marshal(aslMap[key])
	if err != nil {
		fmt.Println(err)
	}
	return string(v)
}