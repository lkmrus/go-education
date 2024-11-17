package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GetPostResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func getPosts(url string) (*[]GetPostResponse, error) {
	if url == "" {
		return nil, fmt.Errorf("url is empty")
	}

	var result, err = http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer result.Body.Close()

	var PostData []GetPostResponse
	err = json.NewDecoder(result.Body).Decode(&PostData)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &PostData, nil
}

func showLoader() {
	//	for with timeout in cycle
	symbols := []string{"|", "/", "-", "\\"}
	for i := 0; i < 1000; i++ {
		for _, val := range symbols {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(val)
		}
	}
}

func main() {
	result, err := getPosts("https://my-json-server.typicode.com/typicode/demo/posts")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	showLoader()
	fmt.Println("Result:")
	fmt.Println(result)
}
