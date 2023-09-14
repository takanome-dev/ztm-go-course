//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func TestAddHealthAndEnergy(t *testing.T) {
	player := CreatePlayer()
	player.addHealth(200)

	if player.health > player.maxHealth {
		t.Fatalf("Health should not go above the limit: %v, want %v", player.health, player.maxHealth)
	}

	player.applyDamage(500)

	if player.health < 0 {
		t.Fatalf("Health should not go below 0: %v", player.health)
	}

	player.addEnergy(1000)

	if player.energy > player.maxEnergy {
		t.Fatalf("Energy should not go above the limit: %v, want %v", player.energy, player.maxEnergy)
	}

	player.consumeEnergy(9999)

	if player.energy < 0 {
		t.Fatalf("Energy should not go below 0: %v", player.energy)
	}
}
