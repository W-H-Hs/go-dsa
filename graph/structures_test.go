package graph

import (
	"fmt"
	"testing"
)

var g = NewGraph(false)

func init() {
	for i := 65; i < 81; i++ {
		_ = g.InsertVertex(string([]byte{byte(i)}))
	}
	_, _ = g.InsertEdge(g.GetVertex("A"), g.GetVertex("B"), 0)
	_, _ = g.InsertEdge(g.GetVertex("B"), g.GetVertex("C"), 0)
	_, _ = g.InsertEdge(g.GetVertex("C"), g.GetVertex("D"), 0)
	_, _ = g.InsertEdge(g.GetVertex("E"), g.GetVertex("F"), 0)
	_, _ = g.InsertEdge(g.GetVertex("A"), g.GetVertex("E"), 0)
	_, _ = g.InsertEdge(g.GetVertex("B"), g.GetVertex("F"), 0)
	_, _ = g.InsertEdge(g.GetVertex("C"), g.GetVertex("G"), 0)
	_, _ = g.InsertEdge(g.GetVertex("D"), g.GetVertex("H"), 0)
	_, _ = g.InsertEdge(g.GetVertex("A"), g.GetVertex("F"), 0)
	_, _ = g.InsertEdge(g.GetVertex("D"), g.GetVertex("G"), 0)
	_, _ = g.InsertEdge(g.GetVertex("I"), g.GetVertex("J"), 0)
	_, _ = g.InsertEdge(g.GetVertex("J"), g.GetVertex("K"), 0)
	_, _ = g.InsertEdge(g.GetVertex("E"), g.GetVertex("I"), 0)
	_, _ = g.InsertEdge(g.GetVertex("G"), g.GetVertex("K"), 0)
	_, _ = g.InsertEdge(g.GetVertex("H"), g.GetVertex("L"), 0)
	_, _ = g.InsertEdge(g.GetVertex("F"), g.GetVertex("I"), 0)
	_, _ = g.InsertEdge(g.GetVertex("G"), g.GetVertex("J"), 0)
	_, _ = g.InsertEdge(g.GetVertex("G"), g.GetVertex("L"), 0)
	_, _ = g.InsertEdge(g.GetVertex("M"), g.GetVertex("N"), 0)
	_, _ = g.InsertEdge(g.GetVertex("I"), g.GetVertex("M"), 0)
	_, _ = g.InsertEdge(g.GetVertex("K"), g.GetVertex("O"), 0)
	_, _ = g.InsertEdge(g.GetVertex("L"), g.GetVertex("P"), 0)
	_, _ = g.InsertEdge(g.GetVertex("I"), g.GetVertex("N"), 0)
	_, _ = g.InsertEdge(g.GetVertex("K"), g.GetVertex("N"), 0)
}

func TestGraph(t *testing.T) {
	g := NewGraph(false)
	fmt.Println(g)
	BFS(g, g.GetVertex("A"), func(vertex *Vertex) bool {
		fmt.Println(vertex)
		return true
	})
}

func TestGraph_GetVertex(t *testing.T) {
	fmt.Println(g.GetVertex("A"))
}

func TestGraph_GetEdge(t *testing.T) {
	// 存在的边。
	fmt.Println(g.GetEdge(g.GetVertex("A"), g.GetVertex("B")))
	// 不存在的边。
	fmt.Println(g.GetEdge(g.GetVertex("F"), g.GetVertex("G")))
}

func TestGraph_GetDegreeOfVertex(t *testing.T) {
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("A"), false))
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("B"), false))
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("C"), false))
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("D"), false))
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("E"), false))
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("F"), false))
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("G"), false))
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("H"), false))
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("I"), false))
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("J"), false))
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("K"), false))
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("L"), false))
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("M"), false))
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("N"), false))
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("O"), false))
	fmt.Println(g.GetDegreeOfVertex(g.GetVertex("P"), false))
}

func TestGraph_GetIncidentEdges(t *testing.T) {
	fmt.Println(g.GetIncidentEdges(g.GetVertex("A"), false))
}

func TestGraph_RemoveVertex(t *testing.T) {
	fmt.Println(g.GetVertexCount())
	_, err := g.RemoveVertex(g.GetVertex("F"))
	if err != nil {
		panic(err)
	}
	fmt.Println(g.GetVertexCount())
	BFS(g, g.GetVertex("A"), func(vertex *Vertex) bool {
		fmt.Println(vertex)
		return true
	})
}

func TestGraph_RemoveEdge(t *testing.T) {
	_, err := g.RemoveEdge(g.GetVertex("A"), g.GetVertex("F"))
	if err != nil {
		panic(err)
	}
	BFS(g, g.GetVertex("A"), func(vertex *Vertex) bool {
		fmt.Println(vertex)
		return true
	})
}
