package main

import "fmt"

func main() {
	var myName = "takanome"
	fmt.Println("my name is", myName)

	var name string = "dev"
	fmt.Println("name =", name)

	username := "admin"
	fmt.Println("username = ", username)

	var sum int
	fmt.Println("The sum is ", sum)

	part1, other := 1, 5
	fmt.Println("part1 is", part1, "and other is", other)

	part2, other := 3, 0
	fmt.Println("part2 is", part2, "and other is", other)

	sum = part1 + part2
	fmt.Println("The actual sum is", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("This lesson is about", lessonName, "and we are doing a", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)
}