package graph

import (
	"errors"
	"go-dsa/list"
	"reflect"
)

// 图中的点结构。
type Vertex struct {
	// 点包含的元素。
	Element interface{}
}

func (v *Vertex) String() string {
	return reflect.ValueOf(v.Element).String()
}

func NewVertex(ele interface{}) *Vertex {
	return &Vertex{Element: ele}
}

// 图中的边结构。
type Edge struct {
	// 边起点。
	origin *Vertex
	// 边终点。
	destination *Vertex
	// 权重。
	Weight int
}

func (e *Edge) String() string {
	return e.origin.String() + "->" + e.destination.String()
}

func NewEdge(origin, destination *Vertex, weight int) *Edge {
	return &Edge{
		origin:      origin,
		destination: destination,
		Weight:      weight,
	}
}

// 获取边起点。
func (e *Edge) GetOrigin() *Vertex {
	return e.origin
}

// 获取边终点。
func (e *Edge) GetDestination() *Vertex {
	return e.destination
}

// 给定边中的一点，返回边的另外一点。
func (e *Edge) Opposite(v *Vertex) *Vertex {
	switch v {
	case e.origin:
		return e.destination
	case e.destination:
		return e.origin
	default:
		return nil
	}
}

type Graph struct {
	// 标识有向图与否。
	isDirected bool

	// 点与边的对应关系。
	outgoing map[*Vertex]map[*Vertex]*Edge
	incoming map[*Vertex]map[*Vertex]*Edge

	// 数组的索引为顶点的唯一标识，该数组主要用在并查集中。
	vertexes list.List
	// 根据权值大小排序的边，如果是无权边，则根据插入顺序排序。
	sortedEdges list.List
}

// 构造函数。
func NewGraph(isDirected bool) *Graph {
	g := &Graph{
		isDirected:  isDirected,
		outgoing:    make(map[*Vertex]map[*Vertex]*Edge),
		vertexes:    list.NewSingleCycleLinkedList(),
		sortedEdges: list.NewSingleCycleLinkedList(),
	}
	if isDirected {
		g.incoming = make(map[*Vertex]map[*Vertex]*Edge)
	}

	return g
}

// 返回图是否是有向图。
func (g *Graph) IsDirected() bool {
	return g.isDirected
}

// 获取图中所有节点的数量。
func (g *Graph) GetVertexCount() int {
	return len(g.outgoing)
}

// 获取图中所有边的数量。
func (g *Graph) GetEdgeCount() int {
	count := 0
	for _, v := range g.outgoing {
		if v != nil {
			count++
		}
	}
	if g.isDirected {
		return count
	} else {
		return count >> 1
	}
}

// 给定元素element，返回图中对应的顶点。
func (g *Graph) GetVertex(element interface{}) *Vertex {
	vertexes := reflect.ValueOf(g.outgoing).MapKeys()
	for i := 0; i < len(vertexes); i++ {
		v := vertexes[i].Interface().(*Vertex)
		if v.Element == element {
			return v
		}
	}

	return nil
}

func (g *Graph) GetAllVertex() []*Vertex {
	var vertexes []*Vertex
	for k := range g.outgoing {
		vertexes = append(vertexes, k)
	}
	return vertexes
}

// 获取从origin到destination的边。
func (g *Graph) GetEdge(origin, destination *Vertex) *Edge {
	return g.outgoing[origin][destination]
}

// 获取给定点的度，有向图中可通过参数isOutgoing控制获取出度或入度。
func (g *Graph) GetDegreeOfVertex(v *Vertex, isOutgoing bool) int {
	if g.isDirected && !isOutgoing {
		return len(g.incoming[v])
	} else {
		return len(g.outgoing[v])
	}
}

// 返回v的所有边，有向图可通过参数isOutgoing控制获取出度边或入度边。
func (g *Graph) GetIncidentEdges(v *Vertex, isOutgoing bool) map[*Vertex]*Edge {
	if g.isDirected && !isOutgoing {
		return g.incoming[v]
	} else {
		return g.outgoing[v]
	}
}

// 插入顶点，返回插入顶点。
func (g *Graph) InsertVertex(element interface{}) *Vertex {
	v := NewVertex(element)

	g.outgoing[v] = make(map[*Vertex]*Edge)
	if g.isDirected {
		g.incoming[v] = make(map[*Vertex]*Edge)
	}

	// 将顶点插入到数组中。
	g.vertexes.Append(v)

	return v
}

// 插入边，返回插入边。
func (g *Graph) InsertEdge(origin, destination *Vertex, weight int) (*Edge, error) {
	if origin == nil || destination == nil {
		return nil, errors.New("origin or destination is nil")
	}

	// 检查边的起点是否在图中。
	if _, ok := g.outgoing[origin]; !ok {
		return nil, errors.New("origin of edge isn't in graph")
	}
	// 检查边的终点是否在图中。
	if _, ok := g.outgoing[destination]; !ok {
		return nil, errors.New("destination of edge isn't in graph")
	}

	e := NewEdge(origin, destination, weight)
	g.outgoing[origin][destination] = e
	if g.isDirected {
		// 如果是有向图，在incoming中添加终点的入度边，此时该边形成有向边，只能从origin到destination。
		g.incoming[destination][origin] = e
	} else {
		// 如果是无向图，在outgoing中添加终点的入度边，此时该边无向边，既能从origin到destination，也能从destination到
		// origin。
		g.outgoing[destination][origin] = e
	}

	// 使用二分法查找e在sortedEdges中的插入位置。
	begin := 0
	end := g.sortedEdges.Len()
	for begin < end {
		midIdx := (begin + end) >> 1
		midEle, _ := g.sortedEdges.Get(midIdx)
		if e.Weight >= midEle.(*Edge).Weight {
			begin = midIdx + 1
		} else {
			end = midIdx
		}
	}
	// 插入e。
	_ = g.sortedEdges.Insert(end, e)

	return e, nil
}

// 删除给定的节点，返回被删除节点。
func (g *Graph) RemoveVertex(v *Vertex) (*Vertex, error) {
	if v == nil {
		return nil, nil
	}

	// 检查给定点是否在图中。
	if _, ok := g.outgoing[v]; !ok {
		return nil, errors.New("vertex isn't in graph")
	}

	// 获取v所有边的另一个点。
	destinations := g.outgoing[v]
	// 删除给定节点。
	delete(g.outgoing, v)
	// 删除另一个点对应的v点。
	for k := range destinations {
		if g.isDirected {
			delete(g.incoming[k], v)
		} else {
			delete(g.outgoing[k], v)
		}
	}

	_, _ = g.vertexes.Remove(g.vertexes.Len() - 1)

	return v, nil
}

// 删除给定节点之间的边，返回被删除边。
func (g *Graph) RemoveEdge(origin, destination *Vertex) (*Edge, error) {
	// 检查边的起点是否在图中。
	if _, ok := g.outgoing[origin]; !ok {
		return nil, errors.New("origin of edge isn't in graph")
	}
	// 检查边的终点是否在图中。
	if _, ok := g.outgoing[destination]; !ok {
		return nil, errors.New("destination of edge isn't in graph")
	}

	e := g.outgoing[origin][destination]
	delete(g.outgoing[origin], destination)
	if g.isDirected {
		delete(g.incoming[destination], origin)
	} else {
		delete(g.outgoing[destination], origin)
	}

	for i := 0; i < g.sortedEdges.Len(); i++ {
		edge, _ := g.sortedEdges.Get(i)
		if edge.(*Edge) == e {
			_, _ = g.sortedEdges.Remove(i)
		}
	}

	return e, nil
}
