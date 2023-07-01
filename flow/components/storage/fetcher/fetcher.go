package fetcher

import "context"

type Fetcher interface {
	Fetch(context.Context, string) string
}