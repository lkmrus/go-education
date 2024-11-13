package main

import "fmt"

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	return a / b
}

func execMath(function MathFunc, a, b float64) {
	mathFunc := function
	if mathFunc == nil {
		fmt.Println("Function not found")
		return
	}

	result := mathFunc(a, b)
	fmt.Println("Result: ", result)
}

type MathFunc func(a, b float64) float64
type Math map[string]MathFunc

var mathMap = Math{
	"add": add,
	"sub": sub,
	"mul": mul,
	"div": div,
}

func main() {
	var funcName string
	fmt.Println("Enter function name: ")
	fmt.Scan(&funcName)

	var math MathFunc = mathMap[funcName]
	if math == nil {
		fmt.Println("Function not found")
		return
	}

	fmt.Println("Enter function args: ")
	var a, b float64
	fmt.Scan(&a, &b)
	result := math(a, b)
	fmt.Println("Result: ", result)
	execMath(mathMap["add"], 1, 2)
}
