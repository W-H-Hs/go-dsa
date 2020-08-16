package graph

import (
	"go-dsa/list"
	"go-dsa/unionfind"
	"reflect"
)

// 广度优先搜索，类似于二叉树的层序遍历。
func BFS(g *Graph, v *Vertex, exec func(*Vertex) bool) {
	if g == nil || v == nil || exec == nil {
		return
	}
	// 使用并查集区分已遍历顶点和未遍历顶点。
	uf := unionfind.NewUnionFind(g.GetAllVertex(), true)
	// BFS的辅助队列。
	queue := list.NewCycleQueue(g.GetDegreeOfVertex(v, g.isDirected))

	queue.EnQueue(v)
	for queue.Len() != 0 {
		vertex := queue.DeQueue().(*Vertex)
		// 将vertex交给调用者处理。
		if !exec(vertex) {
			break
		}
		// 获取vertex的所有出度边的终点。
		destinations := reflect.ValueOf(g.outgoing[vertex]).MapKeys()
		for i := 0; i < len(destinations); i++ {
			destination := destinations[i].Interface().(*Vertex)
			// 如果终点在并查集中所在集合不是已遍历集合（即未遍历），则添加到队列中进行遍历。
			if uf.Find(destination) != uf.Find(v) {
				queue.EnQueue(destinations[i].Interface().(*Vertex))
				// 将vertex合并到并查集已遍历集合中，该集合的根节点为v。
				uf.Union(v, destination)
			}
		}
	}
}

func search(graph *Graph, vertex *Vertex, execute func(*Vertex) bool, u unionfind.UnionFind) {
	destinations := reflect.ValueOf(graph.outgoing[vertex]).MapKeys()
	for i := 0; i < len(destinations); i++ {
		destination := destinations[i].Interface().(*Vertex)
		if u.Find(destination) != u.Find(vertex) {
			execute(destination)
			u.Union(vertex, destination)
			search(graph, destination, execute, u)
		}
	}
}

// 深度优先搜索，类似于二叉树的前序遍历。
func DFS(g *Graph, v *Vertex, exec func(*Vertex) bool) {
	// 使用并查集区分已遍历顶点和未遍历顶点。
	uf := unionfind.NewUnionFind(g.GetAllVertex(), true)
	// 首先遍历给定节点。
	exec(v)
	// 递归遍历。
	search(g, v, exec, uf)
}
