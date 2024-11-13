package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func loopback() func() {
	i := 10
	return func() {
		i += 1
		fmt.Println(i)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	TEST := os.Getenv("TEST")
	envArray := os.Environ()
	for _, env := range envArray {
		fmt.Println(env)
	}
	fmt.Println(TEST)

}
