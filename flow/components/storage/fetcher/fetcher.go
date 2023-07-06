package fetcher

import "context"

type Fetcher interface {
	Fetch(context.Context, map[string]interface{}) string
}