package pathfindings

import (
	"math"
)

// Dijkstra finds the shortest path to all vertices using Dijkstra's Algorithm
func Dijkstra(graph [][]int, src int) []int {
	// TODO: Implementing in Go!
	n := len(graph)            // Graph's size
	dist := make([]int, n)     // Distance matrix
	visited := make([]bool, n) // Manage visited-already vertices
	// Pre-fill distance nodes with infinity value
	for i := 0; i < n; i++ {
		if i != src {
			dist[i] = math.MaxInt
		}
	}
	cur := src  // Define currently selected vertex
	var min int // Declare minimum (shortest) distance
	var minIndex int
	// Exclude final node
	for iter := 0; iter < n; iter++ {
		// Loop through all adjacent (connected) vertex of the current selected node
		for v := 0; v < n; v++ {
			// Skip visited vertices
			if visited[v] {
				continue
			}
			// Relaxation
			if graph[cur][v] > 0 && dist[cur]+graph[cur][v] < dist[v] {
				dist[v] = dist[cur] + graph[cur][v]
			}
		}
		// Compare and select the next optimal vertex
		min = math.MaxInt
		for v := 0; v < n; v++ {
			// Skip visited vertices
			if visited[v] {
				continue
			}
			// Obtain minimum index
			if dist[v] > 0 && dist[v] < min {
				min = dist[v]
				minIndex = v
			}
		}
		// If dead-end (every adjacent nodes are infinity)
		if min == math.MaxInt {
			return dist
		}
		// Mark as visited
		visited[cur] = true
		// Update currently selected index
		cur = minIndex
	}
	return dist
}
