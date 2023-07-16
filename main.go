package main

// import (
// 	"fmt"

// 	"internal/graph"
// )

func main() {
	// Create a new graph
	g := graph.NewGraph()

	// Add users and establish connections
	g.AddUser(1, "John")
	g.AddUser(2, "Alice")
	g.AddUser(3, "Bob")

	g.AddConnection(1, 2)
	g.AddConnection(2, 3)

	// Perform graph operations and analysis
	// ...
}
