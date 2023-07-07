package flowasl

var fetcher Fetcher

func ConfigFetcher(f Fetcher) {
	fetcher = f
}

func GetFetcher() Fetcher {
	return fetcher
}