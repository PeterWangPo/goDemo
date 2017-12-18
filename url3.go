package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	ch := make(chan string)
	urls := os.Args[1:]
	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
}
func fetch(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error", err)
	}
	defer resp.Body.Close()
	ch <- fmt.Sprintf("statusCode: %d", resp.StatusCode)
}
