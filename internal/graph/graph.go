package graph


type User struct {
	ID         int
	Name       string
	Connections []int
	Interests  []string
}

type Graph struct {
	Users map[int]*User
	Visited map[int]bool
}

func NewGraph() *Graph {
	return &Graph{
		Users: make(map[int]*User),
		Visited: make(map[int]bool),
	}
}

func (g *Graph) AddUser(id int, name string) {
	g.Users[id] = &User{
		ID:         id,
		Name:       name,
		Connections: []int{},
		Interests:  []string{},
	}
}


// Use globally query the visited state of the nodes in a graph
func (g *Graph) IsVisited(id int) bool {
	return g.Visited[id]
}

func (g *Graph) AddConnection(userID, connectionID int) {
	user, ok := g.Users[userID]
	if !ok {
		return // Or handle the case when the user is not found
	}

	// Check if the connection already exists
	for _, conn := range user.Connections {
		if conn == connectionID {
			return // connect already there. do nothing and return.
		}
	}

	user.Connections = append(user.Connections, connectionID)

	// Update the connection on the other side as well
	connectedUser, ok := g.Users[connectionID]
	if !ok {
		return // Or handle the case when the user is not found
	}

	// Check if the connection already exists
	for _, conn := range connectedUser.Connections {
		if conn == userID {
			return // connect already there. do nothing and return.
		}
	}
	connectedUser.Connections = append(connectedUser.Connections, userID)

}


// CalculateDegreeCentrality method
func (g *Graph) CalculateDegreeCentrality() map[int]float64 {
	degreeCentrality := make(map[int]float64)

	for id, user := range g.Users {
		degreeCentrality[id] = float64(len(user.Connections))
	}

	return degreeCentrality
}

// FindCommonConnections method
func (g *Graph) FindCommonConnections(userID1, userID2 int) []int {
	commonConnections := []int{}

	user1 := g.Users[userID1]
	user2 := g.Users[userID2]

	for _, conn1 := range user1.Connections {
		for _, conn2 := range user2.Connections {
			if conn1 == conn2 {
				commonConnections = append(commonConnections, conn1)
			}
		}
	}

	return commonConnections
}
