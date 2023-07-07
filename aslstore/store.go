package aslstore

import "github.com/workflow/flow/components/storage/flowasl"

func init() {
	flowasl.ConfigFetcher(&LocalFile{})
}