package main

import (
	"fmt"
	"github.com/oleiade/lane"
)

type PersonInfo struct {
	Name       string
	Level      int
	IsEngineer bool
}

func main() {
	graph := getPeopleGraph()
	queue := lane.NewQueue()
	searched := make(map[string]struct{})

	// I don't want the search to go over me because I'm an engineer myself
	searched["Me"] = struct{}{}

	// Let's start from my immediate network
	for _, person := range graph["Me"] {
		queue.Enqueue(person)
	}

	for !queue.Empty() {
		person := queue.Dequeue().(PersonInfo)
		if _, ok := searched[person.Name]; !ok {
			if person.IsEngineer {
				fmt.Printf("Found the engineer: %s, degree of connection: %d\n\n", person.Name, person.Level)
				break
			}

			for _, person := range graph[person.Name] {
				queue.Enqueue(person)
			}

			searched[person.Name] = struct{}{}
		}
	}
}

// Defining the initial graph as an indirected graph (mutual relationships)
func getPeopleGraph() map[string][]PersonInfo {
	graph := make(map[string][]PersonInfo)

	me := PersonInfo{Name: "Me", Level: 0, IsEngineer: true}
	alice := PersonInfo{Name: "Alice", Level: 1}
	bob := PersonInfo{Name: "Bob", Level: 1}
	claire := PersonInfo{Name: "Claire", Level: 1, IsEngineer: false}
	dean := PersonInfo{Name: "Dean", Level: 2, IsEngineer: true}
	edward := PersonInfo{Name: "Edward", Level: 2}
	frank := PersonInfo{Name: "Frank", Level: 3}
	gordon := PersonInfo{Name: "Gordon", Level: 3, IsEngineer: true}
	helena := PersonInfo{Name: "Helena", Level: 3}

	graph[me.Name] = []PersonInfo{alice, bob, claire}
	graph[alice.Name] = []PersonInfo{me}
	graph[bob.Name] = []PersonInfo{me}
	graph[claire.Name] = []PersonInfo{me, dean, edward}
	graph[dean.Name] = []PersonInfo{claire}
	graph[edward.Name] = []PersonInfo{claire, frank, gordon, helena}
	graph[frank.Name] = []PersonInfo{edward}
	graph[gordon.Name] = []PersonInfo{edward}
	graph[helena.Name] = []PersonInfo{edward}

	return graph
}
