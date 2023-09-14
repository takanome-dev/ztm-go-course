//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

// * Create a structure to store items and their security tag state
//   - Security tags have two states: active (true) and inactive (false)
type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

// * Create functions to activate and deactivate security tags using pointers
func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

// * Create a checkout() function which can deactivate all tags in a slice
func checkout(items []Item) {
	fmt.Println("Checking out...")

	for _, item := range items {
		deactivate(&item.tag)
	}
}

func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	shirts := Item{"Shirt", Active}
	pants := Item{"Pants", Active}
	shoes := Item{"Shoes", Active}
	purse := Item{"Purse", Active}

	//  - Store them in a slice or array
	items := []Item{shirts, pants, shoes, purse}
	fmt.Println(items)

	//  - Deactivate any one security tag in the array/slice
	deactivate(&items[0].tag)
	fmt.Println("Item 0 deactivated")

	//  - Call the checkout() function to deactivate all tags
	checkout(items)
}
