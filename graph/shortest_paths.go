package graph

const infinity = "infinity"

// 单源最短路径。
func Dijkstra(g *Graph, origin *Vertex) (map[*Vertex]interface{}, map[*Vertex][]*Vertex) {
	// TODO: 未优化。
	if g == nil || origin == nil || len(g.outgoing[origin]) == 0 {
		return nil, nil
	}

	// 该map用于记录被选中的顶点。
	selectedVertexes := map[*Vertex]interface{}{origin: nil}
	// 初始化origin到destination的总权值映射。
	weights := map[*Vertex]interface{}{origin: 0}
	for _, destination := range g.GetAllVertex() {
		if destination != origin {
			edge := g.GetEdge(origin, destination)
			if edge != nil {
				weights[destination] = edge.Weight
			} else {
				// 如果edge为nil，代表origin不能通过一条边到达destination，因此将它们之间的权重初始化为无穷大。
				weights[destination] = infinity
			}
		}
	}
	// 初始化路径为graph的一个子图，其中的所有键为已遍历顶点。
	paths := map[*Vertex][]*Vertex{origin: {origin}}
	visitedVertexCount := 1 // 给定的origin默认为已遍历，因此该变量初始化为1。
	// 最后加入已遍历集合的顶点的所有边的另一个顶点一定是被遍历过的，所以这里不能为<=。
	for visitedVertexCount < g.GetVertexCount() {
		var (
			selectedEdge           *Edge
			vertex, oppositeVertex *Vertex
		)
		// 寻找位于已遍历集合的顶点所有出度边中权值最小的边，该边的另一个顶点必须为未遍历。
		// TODO: 可以用最小堆进行优化。
		for v := range selectedVertexes {
			for _, edge := range g.GetIncidentEdges(v, g.isDirected) {
				// 当edge的另一个点未遍历时，执行初始化或寻找权值最小边的操作。
				if _, ok := selectedVertexes[edge.Opposite(v)]; !ok {
					//if uf.Find(edge.Opposite(v)) != origin{
					if selectedEdge == nil {
						// 如果selectedEdge为nil，则初始化。
						selectedEdge = edge
						vertex = v
						oppositeVertex = selectedEdge.Opposite(v)
					} else {
						// 如果selectedEdge不为nil，则寻找origin关联边中的最小值。
						if selectedEdge.Weight+weights[vertex].(int) > edge.Weight+weights[v].(int) {
							selectedEdge = edge
							vertex = v
							oppositeVertex = selectedEdge.Opposite(v)
						}
					}
				}
			}
		}

		if selectedEdge != nil {
			// 当源顶点到选中边某点的总权值加上边的权值小于边的另一个顶点时，更新另一个顶点的总权值及最短路径。
			if weights[oppositeVertex] == infinity ||
				weights[vertex].(int)+selectedEdge.Weight <= weights[oppositeVertex].(int) {
				weights[oppositeVertex] = weights[vertex].(int) + selectedEdge.Weight
				paths[oppositeVertex] = append(paths[vertex], oppositeVertex)
				selectedVertexes[oppositeVertex] = nil
				visitedVertexCount++
			}
			// 松弛操作，更新oppositeVertex点所有关联边另一个顶点的权值和最短路径。
			for _, edge := range g.GetIncidentEdges(oppositeVertex, g.isDirected) {
				if weights[edge.Opposite(oppositeVertex)] == infinity ||
					weights[oppositeVertex].(int)+edge.Weight < weights[edge.Opposite(oppositeVertex)].(int) {
					weights[edge.Opposite(oppositeVertex)] = weights[oppositeVertex].(int) + edge.Weight
					paths[edge.Opposite(oppositeVertex)] = append(paths[oppositeVertex], edge.Opposite(oppositeVertex))
				}
			}
		}
	}

	return weights, paths
}

// 单源最短路径。
func BellmanForm(g *Graph, origin *Vertex) (map[*Vertex]interface{}, map[*Vertex][]*Vertex, bool) {
	if g == nil || origin == nil || len(g.outgoing) == 0 {
		return nil, nil, false
	}

	// 记录origin到各顶点的最小权值。
	weights := make(map[*Vertex]interface{})
	// 记录origin到各顶点的最短路径。
	paths := make(map[*Vertex][]*Vertex)
	// 初始化weights中origin->origin的最小权重为0，最短路径为它自身。
	for v := range g.outgoing {
		if v == origin {
			weights[v] = 0
			paths[v] = []*Vertex{origin}
		} else {
			weights[v] = infinity
		}
	}

	// 标识是否拥有负权环。
	var hasNegativeWeightRing bool
	// 松弛操作函数。
	relax := func(vertexOfEdge *Vertex, weightOfEdge int, needJudge bool) {
		for vertex := range g.outgoing[vertexOfEdge] {
			// 获取vertexOfEdge的所有出度边，计算这些出度边的另一个顶点经过vertexOfEdge的新权值，如果新权值
			// 小于它们的已有权值，则更新。
			newWeight := g.outgoing[vertexOfEdge][vertex].Weight + weights[vertexOfEdge].(int)
			if weights[vertex] == infinity || weights[vertex].(int) > newWeight {
				if needJudge {
					hasNegativeWeightRing = true
					return
				}

				weights[vertex] = newWeight
				paths[vertex] = append(paths[vertexOfEdge], vertex)
			}
		}
	}

	// 标识是否需要判定负权环。
	var needJudge bool
	// 对g的所有边进行V次遍历。
	for k := 0; k < len(g.outgoing)+1; k++ {
		if k == len(g.outgoing) {
			needJudge = true
		}
		for i := 0; i < g.sortedEdges.Len(); i++ {
			edge, _ := g.sortedEdges.Get(i)
			if weights[edge.(*Edge).origin] != infinity {
				// 当edge的源点到给定源点origin的总权值不为无限大时，对它的所有出度边的另一个顶点
				// 执行松弛操作。
				relax(edge.(*Edge).origin, edge.(*Edge).Weight, needJudge)
			}
			if weights[edge.(*Edge).destination] != infinity {
				// 当edge的终点到给定源点origin的总权值不为无限大时，对它的所有出度边的另一个顶点
				// 执行松弛操作。
				relax(edge.(*Edge).destination, edge.(*Edge).Weight, needJudge)
			}
		}
	}
	return weights, paths, hasNegativeWeightRing
}

// 多源最短路径。
func Floyd(g *Graph) (map[*Vertex]map[*Vertex]interface{}, map[*Vertex]map[*Vertex][]*Vertex) {
	weights := make(map[*Vertex]map[*Vertex]interface{})
	paths := make(map[*Vertex]map[*Vertex][]*Vertex)
	allVertex := g.GetAllVertex()

	for _, origin := range allVertex {
		weights[origin] = make(map[*Vertex]interface{})
		paths[origin] = make(map[*Vertex][]*Vertex)
		for _, destination := range allVertex {
			paths[origin][destination] = []*Vertex{origin}
			if origin == destination {
				weights[origin][destination] = 0
			} else {
				if edge, ok := g.outgoing[origin][destination]; ok {
					weights[origin][destination] = edge.Weight
					paths[origin][destination] = append(paths[origin][destination], destination)
				} else {
					weights[origin][destination] = infinity
				}
			}
		}
	}

	for _, transit := range allVertex {
		for _, origin := range allVertex {
			for _, destination := range allVertex {
				weightBetweenOriginAndTransit := weights[origin][transit]
				weightBetweenTransitAndDestination := weights[transit][destination]
				weightsBetweenOriginAndDestination := weights[origin][destination]
				if weightBetweenOriginAndTransit != infinity && weightBetweenTransitAndDestination != infinity {
					newWeight := weightBetweenOriginAndTransit.(int) + weightBetweenTransitAndDestination.(int)
					if weightsBetweenOriginAndDestination == infinity ||
						weightsBetweenOriginAndDestination.(int) > newWeight {
						weights[origin][destination] = newWeight
						paths[origin][destination] = append(
							paths[origin][transit],
							paths[transit][destination][1:]...,
						)
					}
				}
			}
		}
	}

	return weights, paths
}
