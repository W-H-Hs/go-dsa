package unionfind

import "reflect"

type UnionFind interface {
	Find(interface{}) interface{}
	Union(interface{}, interface{})
	IsSame(interface{}, interface{}) bool
}

func NewUnionFind(elements interface{}, isQuickUnion bool) UnionFind {
	// 为elements中的每个元素建立0~len(elements)-1的映射，作为并查集中的元素。
	elementsValue := reflect.ValueOf(elements)
	elementsSlice := elementsValue.Slice(0, elementsValue.Len())

	elementsId := make(map[interface{}]int)
	for i := 0; i < elementsSlice.Len(); i++ {
		elementsId[elementsSlice.Index(i).Interface()] = i
	}
	// 初始化数组，数组的索引代表实体元素，索引对应的数组元素代表其父节点，也是其所在集合的根节点。
	// 在该初始化阶段，索引与数组元素相同，代表实体元素各成一个集合，集合根节点就是它们自身。
	arr := make([]int, len(elementsId))
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	if isQuickUnion {
		return &QuickUnion{elementsId, arr}
	} else {
		return &QuickFind{elementsId, arr}
	}
}
