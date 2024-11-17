package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

func ping(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp.Status)
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

	var wg sync.WaitGroup

	for _, value := range urls {
		fmt.Println(value)
		go func() {
			wg.Add(1)
			ping(value, &wg)
		}()
	}
	wg.Wait()
}
