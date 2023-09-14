//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price int
}

type ShoppingList struct {
	product Product
}

func printInfo(list [4]ShoppingList) {
	sum, totalItems := 0, 0

	for i := 0; i < len(list); i++ {
		item := list[i]
		sum += item.product.price

		if item.product.name != "" {
			totalItems += 1
		}
	}

	fmt.Println("The last item on the list is", list[totalItems-1])
	fmt.Println("The total number of items is", totalItems)
	fmt.Println("The total cost of the items is", sum)
}

func main() {
	shoppingList := [4]ShoppingList{
		{product: Product{name: "Apple", price: 20}},
		{product: Product{name: "Banana", price: 10}},
		{product: Product{name: "Orange", price: 5}},
	}

	printInfo(shoppingList)

	shoppingList[3] = ShoppingList{Product{name: "Bread", price: 30}}
	printInfo(shoppingList)
}
