package graph

type User struct {
	ID         int
	Name       string
	Connections []int
	Interests  []string
}

type Graph struct {
	Users map[int]*User
}

func NewGraph() *Graph {
	return &Graph{
		Users: make(map[int]*User),
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
