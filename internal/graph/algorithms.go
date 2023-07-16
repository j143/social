package graph

import (
	"fmt"
)

func (g *Graph) BFS(startID int) {
	visited := make(map[int]bool)
	queue := []int{startID}

	visited[startID] = true

	for len(queue) > 0 {
		userID := queue[0]
		queue = queue[1:]

		// Process the user with userID
		fmt.Printf("Processing user with ID %d\n", userID)

		for _, connectionID := range g.Users[userID].Connections {
			if !visited[connectionID] {
				visited[connectionID] = true
				queue = append(queue, connectionID)
			}
		}
	}
}

// Implement Dijkstra's algorithm, clustering, influence measurement, and recommendation system as separate functions in this file
