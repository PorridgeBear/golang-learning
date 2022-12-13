package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func getNumber() int {
	return 50
}

// demonstrate receiving a pointer to a string to save on pass by value copy
func showStringAddress(s *string) {
	fmt.Println(s, *s)
}

func main() {
	const locked = "NEVER CHANGE"
	var whole int = 10
	var fraction float32 = 1.23
	var truth bool = true
	var sentence string = "this is a sentence"
	sentence += "."
	sentence += strconv.FormatBool(truth)
	var list [2]string
	list[0] = "one"
	list[1] = "two"
	var num1, num2 int = 20, 30
	var (
		falsy   bool   = false
		stringy string = "fab"
	)
	fromMethod := getNumber()

	fmt.Println(
		locked, whole, fraction, truth, sentence, list,
		reflect.TypeOf(whole), num1, num2, falsy, stringy,
		fromMethod, &sentence,
	)

	showStringAddress(&sentence)
}
