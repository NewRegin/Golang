package main

import (
	"fmt"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, fetcher Fetcher) []string {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Printf("found: %s %q\n", url, body)

	return urls
}

func main() {
	count_c := make(chan int)   //管道：统计string数目
	new_c := make(chan string)  //管道：新抓取的url
	uniq_c := make(chan string) //管道：非重复的url
	dup_c := make(chan string)  //管道：重复的url

	go func() {
		seen := make(map[string]bool)
		for url := range new_c {
			if !seen[url] {
				seen[url] = true
				uniq_c <- url
			} else {
				dup_c <- url
			}
		}
	}()

	new_c <- "http://golang.org/"

	var ( //计算处理完毕的标志
		processed = 0
		found     = 1
	)

	for {
		select {
		case url := <-uniq_c:
			go func(url string, count_c chan int, new_c chan string) {
				urls := Crawl(url, fetcher)
				count_c <- len(urls)
				processed += 1
				for _, ur := range urls {
					new_c <- ur
				}
			}(url, count_c, new_c)

		case <-dup_c:
			processed += 1

		case num := <-count_c:
			found += num

		default:
			time.Sleep(100 * time.Millisecond)
			if processed == found {
				return
			}
		}
	}

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
var fetcher = &fakeFetcher{
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
