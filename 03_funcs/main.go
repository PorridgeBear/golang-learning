package main

import (
	"fmt"
)

func getHello(name string) string {
	return "Hello " + name
}

// passed function
func getHelloWithFn(fn func(s string) string) string {
	return fn("Mr") + " Crossley"
}

// variadic
func summer(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func main() {
	fmt.Println(getHello("ali"))

	// function in a variable
	fn := func(s string) string {
		return "Hola " + s
	}

	fmt.Println(fn("amigo"))

	fmt.Println(getHelloWithFn(fn))

	fmt.Println(summer(10, 20, 30))
}
