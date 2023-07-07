package aslstore

import (
	"context"

	"github.com/workflow/flow/asl/consts"
)

type MemFetcher struct {

}

func (f *MemFetcher) Fetch(ctx context.Context, key string) string {
	return consts.SimpleResultAsl
}