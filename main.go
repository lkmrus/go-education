package main

import "fmt"

func loopback() func() {
	i := 10
	return func() {
		i += 1
		fmt.Println(i)
	}
}

func main() {
	loop := loopback()
	loop()
	loop()
	loop()
	loop()
	loop()
}
