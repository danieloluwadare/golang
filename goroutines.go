package main

import (
	"fmt"
	"net/http"
	"time"
)

type HttpResponse struct {
	latency time.Duration
	url     string
	err     error
}

func get(url string, ch chan<- HttpResponse) {
	start := time.Now()
	_, err := http.Get(url)
	if err != nil {
		ch <- HttpResponse{latency: 0, url: url, err: err}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		httpResponse := HttpResponse{latency: t, url: url, err: nil}
		ch <- httpResponse
	}
}

func InvokeGoRoutine() {
	ch := make(chan HttpResponse)
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.redditxvjs.com",
		"https://www.twitter.com",
	}
	for _, url := range urls {
		go get(url, ch)
	}
	for _, _ = range urls {
		response := <-ch
		if response.err != nil {
			fmt.Printf("Error in GoRoutine: %s => %s error => %v", response.url, response.latency, response.err.Error())
		} else {
			fmt.Printf("Go Routine Url call: %s => %s\n", response.url, response.latency)
		}
	}
}
