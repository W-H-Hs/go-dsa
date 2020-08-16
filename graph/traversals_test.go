package graph

import (
	"fmt"
	"testing"
)

func TestBFS(t *testing.T) {
	BFS(g, g.GetVertex("A"), func(vertex *Vertex) bool {
		fmt.Println(vertex)
		return true
	})
}

func TestDFS(t *testing.T) {
	DFS(g, g.GetVertex("A"), func(vertex *Vertex) bool {
		fmt.Println(vertex)
		return true
	})
}
