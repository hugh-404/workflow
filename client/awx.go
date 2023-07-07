package client

import (
	"context"

	"github.com/pkg/errors"
	"github.com/workflow/flow/asl/consts"
	"github.com/workflow/flow/components/scheduler"
	"github.com/workflow/flow/components/storage/flowasl"
	_ "github.com/workflow/aslstore"
)

type AwxCient struct {}

func (c *AwxCient) Run(ctx context.Context, params map[string]interface{})(interface{}, error) {
	exeCtx, err := (&scheduler.Scheduler{}).StartExecution(ctx, flowasl.GetFetcher().Fetch(ctx, buildKey(params)), params)
	if err != nil {
		return nil, errors.Wrap(err, "Run StartExecution error")
	}
	return exeCtx.GlobalStore[consts.GlobalStore_Result], nil
}

func buildKey(params map[string]interface{}) string {
	return "key1"
}