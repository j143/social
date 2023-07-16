# `Social Network Analysis`

This is a personal project to do Social Network Analysis for learning data structures and algorithms

### Directory Structure

```
social/
├── main.go
├── go.mod
├── go.sum
├── internal/
│   ├── graph/
│   │   ├── graph.go
│   │   └── algorithms.go
│   ├── clustering/
│   │   └── clustering.go
│   ├── influence/
│   │   └── influence.go
│   └── recommendation/
│       └── recommendation.go
└── visualization/
    ├── visualization.go
    └── templates/
        ├── index.html
        └── ...
```

### Explanation of the file structure:

- main.go: The entry point of your application where you can initialize and run the social network analysis tool.
- go.mod and go.sum: Files to manage Go modules and dependencies.
- internal/: A directory for internal packages of your application.
- graph/: Package containing the graph data structure and related algorithms.
- clustering/: Package for implementing clustering algorithms.
- influence/: Package for measuring influence within the social network.
- recommendation/: Package for implementing recommendation algorithms.
- visualization/: Package for visualizing the social network graph and metrics.
- visualization.go: Contains code for handling the visualization of the graph using a library like gonum/graph or go-echarts.
- templates/: Directory to store HTML templates for the visualization.