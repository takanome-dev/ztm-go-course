package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func main() {
	a := []int{2, 3, 5, 1}
	b := []int{1, 2, 7, 4, 9}

	sumOfA := sum(a...)
	fmt.Println("The sum of a is:", sumOfA)

	sumOfB := sum(b...)
	fmt.Println("The sum of b is:", sumOfB)
}
