//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth int
	energy, maxEnergy int
}

func (player *Player) addHealth(amount int) {
	player.health += amount
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}

	fmt.Println(player.name, "add", amount, "percent of health ->", player.health, "percent remaining")
}

func (player *Player) applyDamage(amount int) {
	if (player.health - amount) <= 0 {
		player.health = 0
	} else {
		player.health -= amount
	}

	fmt.Println(player.name, "receive", amount, "percent of damage ->", player.health, "percent of health remaining")
}

func (player *Player) addEnergy(amount int) {
	if (player.energy + amount) > player.maxEnergy {
		player.energy = player.maxEnergy
	} else {
		player.energy += amount
	}

	fmt.Println(player.name, "add", amount, "percent of energy ->", player.energy, "percent remaining")
}

func (player *Player) consumeEnergy(amount int) {
	if (player.energy - amount) <= 0 {
		player.energy = 0
	} else {
		player.energy -= amount
	}

	fmt.Println(player.name, "consume", amount, "percent of energy ->", player.energy, "percent of energy remaining")
}

func CreatePlayer() Player {
	return Player{
		name:      "Knight",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}
}

func main() {
	player := CreatePlayer()
	fmt.Println("Initial player stats:", player)

	player.applyDamage(90)
	player.addHealth(20)
	player.consumeEnergy(300)
	player.addEnergy(150)

	player.consumeEnergy(666)
}
