package graph

type User struct {
	ID         int
	Name       string
	Connections []int
	Interests  []string
}

type Graph struct {
	Users map[int]User
}

func NewGraph() *Graph {
	return &Graph{
		Users: make(map[int]User),
	}
}

func (g *Graph) AddUser(id int, name string) {
	g.Users[id] = User{
		ID:         id,
		Name:       name,
		Connections: []int{},
		Interests:  []string{},
	}
}

func (g *Graph) AddConnection(userID, connectionID int) {
	g.Users[userID].Connections = append(g.Users[userID].Connections, connectionID)
}
