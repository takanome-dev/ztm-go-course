//--Summary:
//  Write a program to display server status.
//
//--Requirements:

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

// * Create a function to print server status, including:
//   - Number of servers
//   - Number of servers for each status (Online, Offline, Maintenance, Retired)
func printStatus(servers map[string]int) {
	fmt.Println("Number of servers:", len(servers))

	statuses := make(map[int]int)

	for _, status := range servers {
		switch status {
		case Online:
			statuses[Online] += 1
		case Offline:
			statuses[Offline] += 1
		case Maintenance:
			statuses[Maintenance] += 1
		case Retired:
			statuses[Retired] += 1
		default:
			panic("unhandled server status")
		}
	}

	fmt.Println(statuses[Online], "servers are online")
	fmt.Println(statuses[Offline], "servers are offline")
	fmt.Println(statuses[Maintenance], "servers are in maintenance")
	fmt.Println(statuses[Retired], "servers are retired")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Store the existing slice of servers in a map
	//* Default all of the servers to `Online`
	serverStatus := make(map[string]int)

	for _, server := range servers {
		serverStatus[server] = Online
	}
	fmt.Println(serverStatus)

	//* Perform the following status changes and display server info:
	//  - display server info
	printStatus(serverStatus)

	//  - change `darkstar` to `Retired`
	serverStatus["darkstar"] = Retired

	//  - change `aiur` to `Offline`
	serverStatus["aiur"] = Offline

	//  - display server info
	printStatus(serverStatus)

	//  - change all servers to `Maintenance`
	for key := range serverStatus {
		serverStatus[key] = Maintenance
	}

	//  - display server info
	printStatus(serverStatus)
}
