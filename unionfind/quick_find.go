package unionfind

type QuickFind struct {
	elementsId map[interface{}]int
	arr        []int
}

// 寻找给定元素所在集合的根节点。
// quick find中元素对应数组的索引，其父节点为数组元素值。
// 且所有元素的父节点即是集合的根节点。
func (qf *QuickFind) Find(element interface{}) interface{} {
	return qf.arr[qf.elementsId[element]]
}

// 给定element1和element2两个元素，合并它们所在的集合。
// v1所在集合的根节点为新集合的根节点。
func (qf *QuickFind) Union(element1, element2 interface{}) {
	// 记录v2所在集合的根节点。
	root := qf.arr[qf.elementsId[element2]]
	// 更改v2的父节点为v1所在集合的根节点。
	qf.arr[qf.elementsId[element2]] = qf.arr[qf.elementsId[element1]]

	// 遍历数组找到所有与v2在同一集合的元素，将它们的根节点改为v1所在集合的根节点。
	for i := 0; i < len(qf.arr); i++ {
		if qf.arr[i] == root {
			qf.arr[i] = qf.arr[qf.elementsId[element1]]
		}
	}
}

// 判断element1和element2是否在同一集合。
func (qf *QuickFind) IsSame(element1, element2 interface{}) bool {
	return qf.arr[qf.elementsId[element1]] == qf.arr[qf.elementsId[element2]]
}
