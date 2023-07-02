package client

import (
	"context"

	"github.com/pkg/errors"
	"github.com/workflow/flow/asl/consts"
	"github.com/workflow/flow/components/scheduler"
	"github.com/workflow/flow/components/storage/fetcher"
)

type AwxCient struct {}

func (c *AwxCient) Run(ctx context.Context, params map[string]interface{})(interface{}, error) {
	fetcher := &fetcher.MemFetcher{}
	exeCtx, err := (&scheduler.Scheduler{}).StartExecution(ctx, fetcher.Fetch(ctx, ""))
	if err != nil {
		return nil, errors.Wrap(err, "Run StartExecution error")
	}
	return exeCtx.GlobalResult[consts.GlobalResult_Result], nil
}