package graph

import (
	"fmt"
	"testing"
)

var graph = NewGraph(false)

func init() {
	graph.InsertVertex("PVD")
	graph.InsertVertex("JFK")
	graph.InsertVertex("BWI")
	graph.InsertVertex("BOS")
	graph.InsertVertex("ORD")
	graph.InsertVertex("DFW")
	graph.InsertVertex("MIA")
	graph.InsertVertex("LAX")
	graph.InsertVertex("SFO")

	_, _ = graph.InsertEdge(graph.GetVertex("BOS"), graph.GetVertex("SFO"), 2704)
	_, _ = graph.InsertEdge(graph.GetVertex("BOS"), graph.GetVertex("ORD"), 867)
	_, _ = graph.InsertEdge(graph.GetVertex("BOS"), graph.GetVertex("JFK"), 187)
	_, _ = graph.InsertEdge(graph.GetVertex("BOS"), graph.GetVertex("MIA"), 1258)
	_, _ = graph.InsertEdge(graph.GetVertex("PVD"), graph.GetVertex("ORD"), 849)
	_, _ = graph.InsertEdge(graph.GetVertex("PVD"), graph.GetVertex("JFK"), 144)
	_, _ = graph.InsertEdge(graph.GetVertex("ORD"), graph.GetVertex("SFO"), 1846)
	_, _ = graph.InsertEdge(graph.GetVertex("ORD"), graph.GetVertex("DFW"), 802)
	_, _ = graph.InsertEdge(graph.GetVertex("ORD"), graph.GetVertex("BWI"), 621)
	_, _ = graph.InsertEdge(graph.GetVertex("ORD"), graph.GetVertex("JFK"), 740)
	_, _ = graph.InsertEdge(graph.GetVertex("JFK"), graph.GetVertex("DFW"), 1391)
	_, _ = graph.InsertEdge(graph.GetVertex("JFK"), graph.GetVertex("BWI"), 184)
	_, _ = graph.InsertEdge(graph.GetVertex("JFK"), graph.GetVertex("MIA"), 1090)
	_, _ = graph.InsertEdge(graph.GetVertex("SFO"), graph.GetVertex("LAX"), 337)
	_, _ = graph.InsertEdge(graph.GetVertex("SFO"), graph.GetVertex("DFW"), 1464)
	_, _ = graph.InsertEdge(graph.GetVertex("DFW"), graph.GetVertex("LAX"), 1235)
	_, _ = graph.InsertEdge(graph.GetVertex("DFW"), graph.GetVertex("MIA"), 1121)
	_, _ = graph.InsertEdge(graph.GetVertex("BWI"), graph.GetVertex("MIA"), 946)
	_, _ = graph.InsertEdge(graph.GetVertex("LAX"), graph.GetVertex("MIA"), 2342)

	for i := 0; i < graph.vertexes.Len(); i++ {
		v, _ := graph.vertexes.Get(i)
		fmt.Println(v.(*Vertex).Element, graph.GetDegreeOfVertex(v.(*Vertex), false))
	}
}

func TestPrim(t *testing.T) {
	edges := Prim(graph)
	for _, ele := range edges {
		fmt.Println(ele.origin, ele.destination, ele.Weight)
	}

	fmt.Println(graph.GetEdge(graph.GetVertex("PVD"), graph.GetVertex("BOS")))
}

func TestKruskal(t *testing.T) {
	edges := Kruskal(graph)
	for _, ele := range edges {
		fmt.Println(ele.origin, ele.destination, ele.Weight)
	}
}
