package main

import (
	"fmt"
	"net/http"
	"io"
	"regexp"
	"sync"
)

//A-Z, a-z, 0-9, -, ., _, ~, :, /, ?, #, [, ], @, !, $, &, ', (, ), *, +, ,, ;, and =
var re_link = regexp.MustCompile(`http[s]*[:][/][/][a-zA-Z0-9-._~:/?#[\]@!$&'()*+,;=]+`)
var visited = make(map[string]int)
var visitedMutex = sync.Mutex{}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type FetchT struct {}
type VisitedError struct {
	url string
}

func (e VisitedError) Error() string {
	return fmt.Sprintf("Already visited url: %s", e.url)
}

func (f FetchT) Fetch(url string) (body string, urls []string, err error){
	visitedMutex.Lock()
	_, found := visited[url]

	if found == true {
		visitedMutex.Unlock()
		return body, urls, VisitedError{url}
	}
	visited[url] = 0
	visitedMutex.Unlock()

	resp, err := http.Get(url)

	if err != nil {
		return body, urls, err
	}

	bodyBytes := make([]byte, 4096)
	var n int
	for err == nil {
		n, err = resp.Body.Read(bodyBytes)
		body += string(bodyBytes[:n])
		if err == io.EOF {
			break
		}
	}

	var uniqUrls = make(map[string]int)
	urlList := re_link.FindAllString(body, -1)

	for i := range urlList {
		uniqUrls[urlList[i]] = 0
	}

	for key := range uniqUrls {
		urls = append(urls, key)
	}

	if err == io.EOF {
		err = nil
	}
	return body, urls, err
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Printf("found: %s %q\n", url, body)
	fmt.Printf("found: %s %d\n", url, len(body))
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	var fetcher FetchT
	Crawl("http://golang.org/", 4, fetcher)
}
