/*
Writing a graph function to find the closest person that is an engineer
*/

package main

import (
	"errors"
	"fmt"
	"github.com/oleiade/lane"
)

type Person struct {
	Name       string
	Level      int
	IsEngineer bool
}

func main() {
	graph := make(map[string][]Person)

	connections, err := getConnections(1, true)
	if err != nil {
		panic(err)
	}

	// Adding my connections - one of them can be an engineer
	graph["me"] = connections

	// Adding second-degree connections to one of my contacts - one of them can be an engineer
	connections2, err := getConnections(2, true)
	if err != nil {
		panic(err)
	}
	// Let's assign those connections to the last person of the previous level
	name := connections[len(connections)-1].Name
	graph[name] = connections2

	// Adding third-degree connections mostly to check if recursion is working as expected
	connections3, err := getConnections(3, true)
	if err != nil {
		panic(err)
	}
	// Let's assign those connections to the last person of the previous level
	name = connections2[len(connections2)-1].Name
	graph[name] = connections3

	queue := enqueueConnections(graph, graph["me"])

	//for !queue.Empty() {
	//	fmt.Println(queue.Dequeue())
	//}

	// Time to find the closest engineer :)
	findClosestEnginneer(queue)

	// Finding all other engineers - this will continue from the point it stopped above
	for !queue.Empty() {
		person := queue.Dequeue().(Person)

		if person.IsEngineer {
			fmt.Printf("Another engineer: %s, degree of connection: %d\n", person.Name, person.Level)
		}
	}
}

// findClosestEngineer executes breadth-first search over a graph
// "breadth-first" just means we want to find the closest connections with a certain characteristic
// first - on our example we are looking for the closest person that happens to be an engineer.
// This is easy after we run enqueueConnections because people are already in proximity order :) we return the first case.
func findClosestEnginneer(queue *lane.Queue) {
	for !queue.Empty() {
		person := queue.Dequeue().(Person)

		if person.IsEngineer {
			fmt.Printf("Found the engineer: %s, degree of connection: %d\n\n", person.Name, person.Level)
			break
		}
	}
}

// enqueueConnections will run recursively over the graph to form the complete map of connections organized in a queue
func enqueueConnections(graph map[string][]Person, people []Person) *lane.Queue {
	queue := lane.NewQueue()
	var deeperConnections []Person

	for _, person := range people {
		queue.Enqueue(person)

		if _, ok := graph[person.Name]; ok {
			deeperConnections = append(deeperConnections, graph[person.Name]...)
		}
	}

	if len(deeperConnections) > 0 {
		q := enqueueConnections(graph, deeperConnections)

		for !q.Empty() {
			queue.Enqueue(q.Dequeue())
		}
	}

	return queue
}

func getConnections(degree int, markAsEngineer bool) ([]Person, error) {
	switch degree {
	case 1:
		return []Person{
			{Name: "Alice", Level: 1},
			{Name: "Bob", Level: 1},
			{Name: "Claire", Level: 1, IsEngineer: markAsEngineer},
		}, nil
	case 2:
		return []Person{
			{Name: "Dean", Level: 2, IsEngineer: markAsEngineer},
			{Name: "Edward", Level: 2},
		}, nil
	case 3:
		return []Person{
			{Name: "Frank", Level: 3},
			{Name: "Gordon", Level: 3, IsEngineer: markAsEngineer},
			{Name: "Helena", Level: 3},
		}, nil

	default:
		return nil, errors.New("this method supports up to 3rd degree connections")
	}
}
