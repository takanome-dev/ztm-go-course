package main

import "fmt"

func double(x int) int {
	return x * 2
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from the greet function!")
}

func main() {
	greet()

	num := 3
	dozen := double(num)
	fmt.Println("The double of", num, "is", dozen)

	num1, num2 := 3, 7
	sum := add(num1, num2)
	fmt.Println("The sum of", num1, "and", num2, "is", sum)
}
