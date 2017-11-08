package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// UrlCounter
type UrlCounter struct {
	urls map[string]bool
	mux  sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (uc UrlCounter) Crawl(url string, depth int, fetcher Fetcher) {
	uc.crawl_sub(url, depth)
	time.Sleep(time.Second)
	return
}

func (uc UrlCounter) crawl_sub(url string, depth int) {
	if depth <= 0 {
		fmt.Printf("too deep: %s\n", url)
		return
	}

	if c := uc.urls[url]; c == true {
		fmt.Printf("skip(already fetched): %s\n", url)
		return
	}
	// Don't fetch the same URL twice.
	uc.mux.Lock()
	uc.urls[url] = true
	uc.mux.Unlock()
	// Fetch URLs in parallel.
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go uc.crawl_sub(u, depth-1)
	}
}

func main() {
	uc := UrlCounter{urls: make(map[string]bool)}
	uc.Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
