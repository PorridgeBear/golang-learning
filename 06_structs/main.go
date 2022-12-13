package main

import (
	"fmt"
	"reflect"
)

type Recipe struct {
	Name        string
	Ingredients []IngredientAmount
}

type IngredientAmount struct {
	Ingredient  Ingredient
	AmountGrams float32
}

type Ingredient struct {
	Name string
}

func main() {

	macaroni := Ingredient{Name: "Macaroni"} // returns an Ingredient

	cheese := new(Ingredient) // returns a pointer to the Ingredient
	cheese.Name = "Cheese"

	theMacaroni := IngredientAmount{Ingredient: macaroni, AmountGrams: 100}
	theCheese := IngredientAmount{Ingredient: *cheese, AmountGrams: 60} // note the pointer dereference
	ingredientAmounts := make([]IngredientAmount, 0)                    // interesting if set to 1 created default instance
	ingredientAmounts = append(ingredientAmounts, theMacaroni, theCheese)

	macAndCheese := Recipe{Name: "Mac & Cheese", Ingredients: ingredientAmounts}
	fmt.Println(reflect.TypeOf(macAndCheese))
	fmt.Println(macAndCheese)
	fmt.Printf("%+v\n", macAndCheese)
}
