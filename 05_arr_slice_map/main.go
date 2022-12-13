package main

import (
	"fmt"
)

func main() {

	// arrays
	var numbers [2]string
	numbers[0] = "one"
	numbers[1] = "two"
	fmt.Println(numbers)

	// slices
	stretchyNumbers := make([]int, 2)
	stretchyNumbers[0] = 1
	stretchyNumbers[1] = 2
	stretchyNumbers = append(stretchyNumbers, 3, 4)
	fmt.Println(stretchyNumbers)

	stretchyNumbers = append(stretchyNumbers[:2], stretchyNumbers[3:]...)
	fmt.Println(stretchyNumbers)

	// maps
	priceList := make(map[string]float32)
	priceList["fish"] = 12.50
	priceList["chips"] = 4.50
	priceList["peas"] = 1
	fmt.Println(priceList)

	delete(priceList, "peas")
	fmt.Println(priceList)
}
