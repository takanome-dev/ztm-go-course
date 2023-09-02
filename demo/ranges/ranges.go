package main

import "fmt"

func main() {
	slice := []string{"Hello", "world", "!"}

	for i, element := range slice {
		fmt.Println("index:", i, "element:", element)

		for _, char := range element {
			fmt.Printf("  %q\n", char)
		}
	}
}
