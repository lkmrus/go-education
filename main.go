package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ping(url string, statusCh chan int, chanErr chan error) {
	resp, err := http.Get(url)
	if err != nil {
		chanErr <- err
	} else {
		statusCh <- resp.StatusCode
	}
}

func main() {
	path := flag.String("file", "sites.txt", "PATH to file")
	flag.Parse()

	result, err := os.ReadFile(*path)
	if err != nil {
		panic(err.Error())
	}

	urls := strings.Split(string(result), "\n")
	statusCh := make(chan int)
	errorCh := make(chan error)
	for _, value := range urls {
		fmt.Println(value)
		go func() {
			ping(value, statusCh, errorCh)
		}()
	}

	for range urls {
		select {
		case status := <-statusCh:
			fmt.Println("Response -> ", status)
		case err := <-errorCh:
			fmt.Println(err)
		}
	}
}
