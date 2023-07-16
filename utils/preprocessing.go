package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID          int
	Connections []int
}

func preprocessDataset(edgesPath string, circlesPath string) (map[int]*User, error) {
	graph := make(map[int]*User)

	// Read edges file
	err := readEdgesFile(edgesPath, graph)
	if err != nil {
		return nil, err
	}

	// Read circles file
	err = readCirclesFile(circlesPath, graph)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func readEdgesFile(filePath string, graph map[int]*User) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		user1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Printf("Invalid user ID: %s\n", parts[0])
			continue
		}

		user2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Printf("Invalid user ID: %s\n", parts[1])
			continue
		}

		addConnection(graph, user1, user2)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func readCirclesFile(filePath string, graph map[int]*User) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")

		// circle := parts[0]
		users := parts[1:]

		for _, userStr := range users {
			userID, err := strconv.Atoi(userStr)
			if err != nil {
				log.Printf("Invalid user ID: %s\n", userStr)
				continue
			}

			if user, ok := graph[userID]; ok {
				user.Connections = append(user.Connections, -1) // Represent membership in a circle with -1
			} else {
				log.Printf("User ID %d not found in the graph\n", userID)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func addConnection(graph map[int]*User, userID1, userID2 int) {
	user1, ok := graph[userID1]
	if !ok {
		user1 = &User{ID: userID1, Connections: []int{}}
		graph[userID1] = user1
	}

	user2, ok := graph[userID2]
	if !ok {
		user2 = &User{ID: userID2, Connections: []int{}}
		graph[userID2] = user2
	}

	user1.Connections = append(user1.Connections, userID2)
	user2.Connections = append(user2.Connections, userID1)
}

func main() {
	edgesPath := "data/facebook/0.edges"
	circlesPath := "data/facebook/0.circles"

	graph, err := preprocessDataset(edgesPath, circlesPath)
	if err != nil {
		log.Fatal(err)
	}

	// printLoadedData(graph)

	// Run the test functions
	testUserConnections(graph)
	testCircleMembership(graph)

	// Perform further analysis or tests
}


func testUserConnections(graph map[int]*User) {
	// Test case: Check if user 0 has connections to users 1, 2, and 3
	user190Connections := graph[190].Connections
	expectedConnections := []int{96, 23, 23, 96, 229, 229, -1}

	if !slicesEqual(user190Connections, expectedConnections) {
		fmt.Println("User 190 connections test failed!")
	} else {
		fmt.Println("User 190 connections test passed!")
	}

	// Add more test cases as needed
}

func slicesEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}

	return true
}


func testCircleMembership(graph map[int]*User) {
	// Test case: Check if user 28 has connection to user 312 and circle membership (-1)
	if !contains(graph[28].Connections, 312) || !contains(graph[28].Connections, -1) {
		fmt.Println("User 28, connections test failed!")
	} else {
		fmt.Println("User 28, connections test passed!")
	}

	// Add more test cases as needed
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// verify the loaded data by printing, look like this:
//
// User ID: 28
// Connections: [312 310 312 144 115 151 226 326 326 226 14 144 140 115 41 41 116 14 149 151 116 310 149 140 -1]
//
func printLoadedData(graph map[int]*User) {
	for userID, user := range graph {
		fmt.Printf("User ID: %d\n", userID)
		fmt.Printf("Connections: %v\n", user.Connections)
	}
}
