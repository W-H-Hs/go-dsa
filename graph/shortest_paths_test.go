package graph

import (
	"fmt"
	"testing"
)

func TestDijkstra(t *testing.T) {
	weights, paths := Dijkstra(graph, graph.GetVertex("BWI"))
	fmt.Println(weights)
	fmt.Println(paths)
}

func TestBellmanForm(t *testing.T) {
	weights, paths, hasNegativeRing := BellmanForm(graph, graph.GetVertex("BWI"))
	fmt.Println(weights)
	fmt.Println(paths)
	fmt.Println(hasNegativeRing)
}

func TestFloyd(t *testing.T) {
	weights, paths := Floyd(graph)
	fmt.Println(weights)
	fmt.Println(paths)
}
