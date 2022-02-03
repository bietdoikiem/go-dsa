package pathfindings

import (
	"fmt"
	"testing"
)

func Test_dijkstra(t *testing.T) {
	graph := [][]int{
		{0, 7, 12, 0, 0, 0}, // A
		{0, 0, 2, 9, 0, 0},  // B
		{0, 0, 0, 0, 10, 0}, // C
		{0, 0, 0, 0, 0, 1},  // D
		{0, 0, 0, 4, 0, 5},  // E
		{0, 0, 0, 0, 0, 0},  // F
	}
	dist := Dijkstra(graph, 0)
	fmt.Println(dist)
}
