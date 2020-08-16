package unionfind

type QuickUnion struct {
	elementsId map[interface{}]int
	arr        []int
}

// 寻找给定元素所在集合的根节点。
// 从v开始找它的父节点，一直到索引对应的数组元素与索引相等，即元素的父节点为元素本身。
// 路径减半优化，使查找路径上每隔一个节点就指向其祖父节点。
func (qu *QuickUnion) Find(element interface{}) interface{} {
	// 初始化孙子节点为给定节点。
	grandson := qu.elementsId[element]
	// 初始化集合根节点为给定节点的父节点。
	root := qu.arr[qu.elementsId[element]]
	// 初始化当前高度为0。
	curr := 0
	for root != qu.arr[root] {
		root = qu.arr[root]
		if curr&1 == 0 {
			// 当更新grandson后遇到第一个偶数高度时，代表找到了grandson的grandpa节点。
			qu.arr[grandson] = root
			grandson = root
		}
		curr++
	}
	// 寻找root对应的元素。
	for k, v := range qu.elementsId {
		if v == root {
			return k
		}
	}
	return nil
}

// 给定element1和element2两个元素，合并它们所在的集合。
// 基于rank的优化，rank高的集合的根节点为新集合的根节点。
func (qu *QuickUnion) Union(element1, element2 interface{}) {
	// 找到v1所在集合的根节点。
	rank1 := 0
	root1 := qu.arr[qu.elementsId[element1]]
	for root1 != qu.arr[root1] {
		root1 = qu.arr[root1]
		rank1++
	}
	// 找到v2所在集合的根节点。
	rank2 := 0
	root2 := qu.arr[qu.elementsId[element2]]
	for root2 != qu.arr[root2] {
		root2 = qu.arr[root2]
		rank2++
	}
	// 合并两个集合。
	if rank1 < rank2 {
		qu.arr[root1] = root2
	} else {
		qu.arr[root2] = root1
	}
}

// 判断element1和element2是否在同一集合。
func (qu *QuickUnion) IsSame(element1, element2 interface{}) bool {
	return qu.Find(element1) == qu.Find(element2)
}
