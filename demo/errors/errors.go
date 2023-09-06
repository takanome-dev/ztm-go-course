package main

import (
	"errors"
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(idx int) (int, error) {
	if idx > len(s.values) {
		return 0, errors.New(fmt.Sprintf("no element at index %v", idx))
	}

	return s.values[idx], nil
}

func main() {
	stuff := Stuff{}

	value, err := stuff.Get(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value is", value)
	}
}
