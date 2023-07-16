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
// func (g *Graph) Dijkstra(startID, targetID int) []int {
//     distances := make(map[int]int)
//     previous := make(map[int]int)
//     pq := make(PriorityQueue, 0)

//     for userID := range g.Users {
//         distances[userID] = math.MaxInt32
//         previous[userID] = -1
//     }

//     distances[startID] = 0
//     pq.Push(&Item{ID: startID, Priority: 0})

//     for pq.Len() > 0 {
//         current := pq.Pop().(*Item)

//         if current.ID == targetID {
//             break
//         }

//         for _, connectionID := range g.Users[current.ID].Connections {
//             weight := 1 // Assuming equal weight for all connections
//             alt := distances[current.ID] + weight

//             if alt < distances[connectionID] {
//                 distances[connectionID] = alt
//                 previous[connectionID] = current.ID
//                 pq.Push(&Item{ID: connectionID, Priority: alt})
//             }
//         }
//     }

//     // Retrieve the shortest path from startID to targetID
//     path := []int{}
//     currentID := targetID

//     for currentID != -1 {
//         path = append([]int{currentID}, path...)
//         currentID = previous[currentID]
//     }

//     return path
// }
