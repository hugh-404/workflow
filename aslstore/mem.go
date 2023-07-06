package aslstore

import (
	"context"

	"github.com/workflow/flow/asl/consts"
)

type MemFetcher struct {

}

func (f *MemFetcher) Fetch(ctx context.Context, keyMap map[string]interface{}) string {
	return consts.SimpleResultAsl
}