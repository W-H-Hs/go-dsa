package graph

import (
	"go-dsa/list"
	"go-dsa/unionfind"
)

// Prim算法：1.从图中任选一顶点作为起始，将其加入已遍历集合中，剩下的顶点作为未遍历集合。
//	  	    2.已遍历集合中所有顶点和未遍历集合中所有顶点之间的边为横切边，从横切边中选出
//            权值最小的边加入最小支撑树边中（如果有多条边都是最小权值，则随机选一条）并从
//            横切边中删除。将权值最小的边的另一个顶点加入到已遍历集合中，同时将该顶点的所
//            有关联边加入横切边中。
//          3.重复步骤2直到所有的顶点都已经遍历。
// Kruskal算法：1.从图中选择权值最小的边和次小的边加入最小支撑树中（这两条边一定在最小支撑
//				 树中），并将这两条边的四个顶点在并查集中加入已遍历集合中。
// 			   2.从除开已被选择边的边集中选择最小的，检测与已被选择的边是否会形成环（从第
//		   	     三小的边就可能会形成环），检测方法是查看该边的两个顶点是否同时在一个已遍
//		   	     历集合中。如果不形成环，则将该边加入最小支撑树中，并将它的两个顶点加入已
//		   	     遍历集合中
//			   3.重复步骤2直到已遍历集合中的点等于图的点。

// Prim算法寻找图g中的最小支撑树。
func Prim(g *Graph) []*Edge {
	if g == nil {
		return nil
	}

	// g中的最小支撑树。
	var edgesOfMinimumSpanningTree []*Edge
	// 并查集用于标识是否在已遍历集合。
	uf := unionfind.NewUnionFind(g.GetAllVertex(), true)
	// 横切边，已遍历集合和未遍历集合之间的所有边。
	crossingEdges := list.NewArrayList()
	// 将第一个添加到图g的顶点的关联边添加到横切边中。
	firstInsertVertex, _ := g.vertexes.Get(0)
	incidentEdges := g.GetIncidentEdges(firstInsertVertex.(*Vertex), false)
	for _, v := range incidentEdges {
		crossingEdges.Append(v)
	}
	// 寻找最小支撑树。
	visitedVertexCount := 2
	for visitedVertexCount <= g.GetVertexCount() {
		// 初始化最小权值边在crossingEdges中的索引为0。
		IdxOfMinimumWeightEdgeOfCrossingEdges := 0
		// 初始化最小权值边为crossingEdge中的0号元素。
		minimumWeightEdge, _ := crossingEdges.Get(IdxOfMinimumWeightEdgeOfCrossingEdges)
		// 寻找crossingEdge中的最小权值边。
		for i := 1; i < crossingEdges.Len(); i++ {
			edge, _ := crossingEdges.Get(i)
			if minimumWeightEdge.(*Edge).Weight > edge.(*Edge).Weight {
				// 如果edge的两个端点不同时在已遍历集合中，则代表它们是横切边（不会出现同时在未遍历集合的情况）。
				minimumWeightEdge = edge
				IdxOfMinimumWeightEdgeOfCrossingEdges = i
			}
		}
		// 将最小权值边添加到边集合中。
		edgesOfMinimumSpanningTree = append(edgesOfMinimumSpanningTree, minimumWeightEdge.(*Edge))
		// 删除crossingEdges中的最小权值边。
		_, _ = crossingEdges.Remove(IdxOfMinimumWeightEdgeOfCrossingEdges)
		// 找到minimumWeightEdge中位于未遍历集合中的点。
		var oppositeVertex *Vertex
		if uf.Find(minimumWeightEdge.(*Edge).origin) !=
			uf.Find(firstInsertVertex.(*Vertex)) {
			oppositeVertex = minimumWeightEdge.(*Edge).origin
		} else {
			oppositeVertex = minimumWeightEdge.(*Edge).destination
		}
		// 将该点合并到已遍历集合中。
		uf.Union(firstInsertVertex.(*Vertex), oppositeVertex)
		// 遍历横切边集合，将两个端点都属于已遍历集合的边移除。
		for i := crossingEdges.Len() - 1; i >= 0; i-- {
			edge, _ := crossingEdges.Get(i)
			if uf.Find(edge.(*Edge).origin) == uf.Find(firstInsertVertex.(*Vertex)) &&
				uf.Find(edge.(*Edge).destination) == uf.Find(firstInsertVertex.(*Vertex)) {
				_, _ = crossingEdges.Remove(i)
			}
		}
		// 找到所有与该点关联的边，并添加到crossingEdges中。
		incidentEdges = g.GetIncidentEdges(oppositeVertex, false)
		for _, v := range incidentEdges {
			// 如果关联边的另一个点不在已遍历集合中，则将其添加进横切边中。
			if uf.Find(v.Opposite(oppositeVertex)) !=
				uf.Find(firstInsertVertex.(*Vertex)) {
				crossingEdges.Append(v)
			}
		}

		visitedVertexCount++
	}

	return edgesOfMinimumSpanningTree
}

func Kruskal(g *Graph) []*Edge {
	if g == nil {
		return nil
	}

	// g中的最小支撑树。
	var edgesOfMinimumSpanningTree []*Edge
	// 并查集用于标识点是否在已遍历集合。
	uf := unionfind.NewUnionFind(g.GetAllVertex(), true)

	// 获取权值最小和次小的边加入最小支撑树中（它们一定包含在最小支撑树中）。
	minimum, _ := g.sortedEdges.Get(0)
	secondSmallest, _ := g.sortedEdges.Get(1)
	edgesOfMinimumSpanningTree = append(edgesOfMinimumSpanningTree, minimum.(*Edge))
	edgesOfMinimumSpanningTree = append(edgesOfMinimumSpanningTree, secondSmallest.(*Edge))

	var visitedVertexCount int
	if minimum.(*Edge).origin != secondSmallest.(*Edge).origin &&
		minimum.(*Edge).origin != secondSmallest.(*Edge).destination &&
		minimum.(*Edge).destination != secondSmallest.(*Edge).origin &&
		minimum.(*Edge).destination != secondSmallest.(*Edge).destination {
		visitedVertexCount = 4

		// 如果这两条边没有公共顶点，则它们自己的两个顶点分别做成两个已遍历集合。
		uf.Union(minimum.(*Edge).origin, minimum.(*Edge).destination)
		uf.Union(secondSmallest.(*Edge).origin, secondSmallest.(*Edge).destination)
	} else {
		visitedVertexCount = 3

		// 如果这两条边有公共顶点，则它们的三个顶点做成一个已遍历集合。
		uf.Union(minimum.(*Edge).origin, minimum.(*Edge).destination)
		uf.Union(minimum.(*Edge).origin, secondSmallest.(*Edge).origin)
		uf.Union(minimum.(*Edge).origin, secondSmallest.(*Edge).destination)
	}
	edgeNum := 2
	for visitedVertexCount <= g.GetVertexCount() {
		edge, _ := g.sortedEdges.Get(edgeNum)
		// 检查edge是否会与已选中边形成环。
		rootOfOrigin := uf.Find(edge.(*Edge).origin)
		rootOfDestination := uf.Find(edge.(*Edge).destination)
		// 如果origin所在集合的根不等于destination所在集合，则代表这条边加入最小支撑树后不会形成环。
		if rootOfOrigin != rootOfDestination {
			edgesOfMinimumSpanningTree = append(edgesOfMinimumSpanningTree, edge.(*Edge))
			if edge.(*Edge).origin == rootOfOrigin {
				if edge.(*Edge).destination == rootOfDestination { // 该边的起点和终点未被遍历过。
					// 将该边的起点和终点合并为一个已遍历集合。
					uf.Union(edge.(*Edge).origin, edge.(*Edge).destination)
					visitedVertexCount += 2
				} else { // 该边的起点未被遍历过，终点被遍历过。
					// 将该边的起点合并到终点所在的已遍历集合。
					uf.Union(edge.(*Edge).destination, edge.(*Edge).origin)
					visitedVertexCount++
				}
			} else {
				// 该边的起点被遍历过，因为有rootOfOrigin!=rootOfDestination的条件语句即该边的起点和终点不可能
				// 在同一集合，因此执行到这里时该边的终点只会出现未被遍历的情况。

				// 将该边的终点合并到起点所在的已遍历集合中。
				uf.Union(edge.(*Edge).origin, edge.(*Edge).destination)
				visitedVertexCount++
			}
		}
		edgeNum++
	}

	return edgesOfMinimumSpanningTree
}
