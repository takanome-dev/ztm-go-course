//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int
type Car string
type Motorcycles string
type Trucks string

type LiftPicker interface {
	PickLift() Lift
}

func (m Motorcycles) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}

func (t Trucks) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}

func (m Motorcycles) PickLift() Lift {
	return SmallLift
}

func (t Trucks) PickLift() Lift {
	return LargeLift
}

func (c Car) PickLift() Lift {
	return StandardLift
}

func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Printf("send %v to small lift\n", p)
	case StandardLift:
		fmt.Printf("send %v to standard lift\n", p)
	case LargeLift:
		fmt.Printf("send %v to large lift\n", p)
	}
}

func main() {
	car := Car("Sporty")
	truck := Trucks("Truck")
	motorcycle := Motorcycles("Croozer")

	sendToLift(motorcycle)
	sendToLift(car)
	sendToLift(truck)
}
