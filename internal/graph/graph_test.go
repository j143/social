package graph

import (
	"reflect"
	"testing"
)

func TestGraph_AddUser(t *testing.T) {
	g := NewGraph()

	g.AddUser(1, "John")
	g.AddUser(2, "Alice")

	if len(g.Users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(g.Users))
	}

	if _, ok := g.Users[1]; !ok {
		t.Error("User with ID 1 not found")
	}

	if _, ok := g.Users[2]; !ok {
		t.Error("User with ID 2 not found")
	}
}

func TestGraph_AddConnection(t *testing.T) {
	g := NewGraph()

	g.AddUser(1, "John")
	g.AddUser(2, "Alice")
	g.AddUser(3, "Bob")

	g.AddConnection(1, 2)
	g.AddConnection(2, 1) // add reverse connection, since connections are bidirectional

	if len(g.Users[1].Connections) != 1 {
		t.Errorf("Expected 1 connection for user 1, got %d", len(g.Users[1].Connections))
	}

	if g.Users[1].Connections[0] != 2 {
		t.Errorf("Expected connection 2 for user 1, got %d", g.Users[1].Connections[0])
	}

	if len(g.Users[2].Connections) != 1 {
		t.Errorf("Expected 1 connection for user 2, got %d", len(g.Users[2].Connections))
	}

	if g.Users[2].Connections[0] != 1 {
		t.Errorf("Expected connection 1 for user 2, got %d", g.Users[2].Connections[0])
	}
}

/*
User 1 (John) --> User 2 (Alice)
  ^               |
  |               v
User 4 (Eve) <-- User 3 (Bob)


*/

func TestGraph_BFS(t *testing.T) {
	g := NewGraph()

	g.AddUser(1, "John")
	g.AddUser(2, "Alice")
	g.AddUser(3, "Bob")
	g.AddUser(4, "Eve")

	g.AddConnection(1, 2)
	g.AddConnection(2, 3)
	g.AddConnection(3, 4)
	g.AddConnection(4, 1)

	// visited := make(map[int]bool)

	g.BFS(1)

	expectedVisited := map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true,
	}

	for id := range g.Users {
		if g.IsVisited(id) != expectedVisited[id] {
			t.Errorf("Expected user with ID %d to be visited, but it was not", id)
		}
	}
}


func TestGraph_CalculateDegreeCentrality(t *testing.T) {
	g := NewGraph()

	g.AddUser(1, "John")
	g.AddUser(2, "Alice")
	g.AddUser(3, "Bob")
	g.AddUser(4, "Eve")

	g.AddConnection(1, 2)
	g.AddConnection(1, 3)
	g.AddConnection(2, 3)
	g.AddConnection(3, 4)

	expectedCentrality := map[int]float64{
		1: 2,
		2: 2,
		3: 3,
		4: 1,
	}

	centrality := g.CalculateDegreeCentrality()

	for id, expected := range expectedCentrality {
		if centrality[id] != expected {
			t.Errorf("Unexpected degree centrality for user %d. Expected: %f, Got: %f", id, expected, centrality[id])
		}
	}
}

func TestGraph_FindCommonConnections(t *testing.T) {
	g := NewGraph()

	g.AddUser(1, "John")
	g.AddUser(2, "Alice")
	g.AddUser(3, "Bob")
	g.AddUser(4, "Eve")

	g.AddConnection(1, 2)
	g.AddConnection(1, 3)
	g.AddConnection(2, 3)
	g.AddConnection(3, 4)

	expectedCommonConnections := []int{3}

	commonConnections := g.FindCommonConnections(1, 2)

	if !reflect.DeepEqual(commonConnections, expectedCommonConnections) {
		t.Errorf("Unexpected common connections between User 1 and User 2. Expected: %v, Got: %v", expectedCommonConnections, commonConnections)
	}
}


