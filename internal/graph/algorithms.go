package graph

import (
	"fmt"
)

func (g *Graph) BFS(startID int) {
	visited := make(map[int]bool)
	queue := make([][2]int, 0) // [0]: user ID, [1]: level
	queue = append(queue, [2]int{startID, 0})

	visited[startID] = true

	// directly modify the Visited map of the Graph struct
	g.Visited = visited

	for len(queue) > 0 {
		userInfo := queue[0]
		queue = queue[1:] // remove the 0th index item from queue

		userID := userInfo[0]
		level := userInfo[1]

		// visited[startID] = true

		// Process the user with userID
		fmt.Printf("User %d (Visited: %v)\n", userID, g.Visited[userID])

		for _, connectionID := range g.Users[userID].Connections {
			if !g.Visited[connectionID] {
				queue = append(queue, [2]int{connectionID, level + 1})
				g.Visited[connectionID] = true
			}
		}
	}
}

// Implement Dijkstra's algorithm, clustering, influence measurement, and recommendation system as separate functions in this file
