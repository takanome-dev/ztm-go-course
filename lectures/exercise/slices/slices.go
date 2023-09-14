//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

// * Create a function to print out the contents of the assembly line
func showLine(line []Part) {
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
}

func main() {
	//* Using a slice, create an assembly line that contains type Part
	assemblyLine := make([]Part, 3)

	//  - Create an assembly line having any three parts
	assemblyLine[0] = "Pipe"
	assemblyLine[1] = "Bolt"
	assemblyLine[2] = "Sheet"

	fmt.Println("3 parts:")
	showLine(assemblyLine)

	//  - Add two new parts to the line
	assemblyLine = append(assemblyLine, "Washer", "Wheel")
	fmt.Println("\nAdded 2 new parts:")
	showLine(assemblyLine)

	//  - Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	fmt.Println("\nSliced so there 2 only the two new lines")
	showLine(assemblyLine)
}
