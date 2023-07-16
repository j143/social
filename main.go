package main

import (
	"fmt"
	"log"

	"./internal/graph"
	"./utils"
)

func main() {
	// Set the file paths for the dataset
	edgesPath := "data/facebook/0.edges"
	circlesPath := "data/facebook/0.circles"

	// Load and preprocess the dataset
	graphData, err := utils.PreprocessDataset(edgesPath, circlesPath)
	if err != nil {
		log.Fatal(err)
	}

	// Perform common connections analysis
	commonConnections, err := graph.FindCommonConnections(userID1, userID2 int)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Common Connections:", commonConnections)

	// Calculate degree centrality
	graph.CalculateDegreeCentrality(graphData)

	// Print degree centrality for each user
	for userID, user := range graphData {
		fmt.Printf("User ID: %d, Degree Centrality: %.2f\n", userID, user.DegreeCentrality)
	}

	// Additional analysis or visualization code can be added here

	// ...
}
