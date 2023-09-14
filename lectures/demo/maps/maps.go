package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)

	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted, new list:", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	if !found {
		fmt.Println("Cereal is not found in list, adding it...")
		shoppingList["cereal"] = cereal
	}

	fmt.Println("Added cereal to list:", shoppingList)

	totalItems := 0

	for _, value := range shoppingList {
		totalItems += value
	}

	fmt.Println("The total items to purchase is", totalItems)
}
