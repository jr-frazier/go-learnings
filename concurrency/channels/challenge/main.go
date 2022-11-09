package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("%s -> error: %s", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	out <- fmt.Sprintf("%s -> %s", url, ctype)
}

// concurrent version
func siteConcurrent(urls []string) {
	ch := make(chan string, len(urls))

	for _, url := range urls {
		ch <- url
	}

	for i := range ch {
		returnType(i, ch)
	}
	close(ch)
}

func main() {

	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}

	// Create response channel
	ch := make(chan string)
	for _, url := range urls {
		go returnType(url, ch)
	}

	for range urls {
		// Run number of URLs times
		out := <-ch
		fmt.Printf(out)
	}

}
