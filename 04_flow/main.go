package main

import (
	"fmt"
)

func main() {

	defer fmt.Println("do this at the end!")
	defer fmt.Println("but this very last")

	truthy := false
	if truthy {
		fmt.Println("so true!")
	}

	balance := 10
	if balance < 0 {
		fmt.Println("overdrawn")
	} else if balance < 10 {
		fmt.Println("soon overdrawn")
	} else {
		fmt.Println("still some life in the account")
	}

	for balance > 0 {
		fmt.Printf("still got dosh %v\n", balance)
		balance -= 5
	}

	op := "ADD"
	switch op {
	case "ADD":
		fmt.Println("adding")
	default:
		fmt.Println("noop")
	}

	for i := 2; i > 0; i-- {
		fmt.Println(i)
	}

	listing := []int{1, 2, 3}
	for index, value := range listing {
		fmt.Printf("index %v value %v\n", index, value)
	}
}
