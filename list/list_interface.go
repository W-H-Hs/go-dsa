package list

type List interface {
	// 返回元素数量
	Len() int
	// 返回是否为空
	IsEmpty() bool
	// 返回是否包含某个元素
	IsContain(interface{}) bool
	// 添加元素到最后面
	Append(interface{})
	// 返回index位置对应的元素
	Get(int) (interface{}, error)
	// 替换index位置的元素，返回被替换的元素
	Set(int, interface{}) (interface{}, error)
	// 向index位置添加元素
	Insert(int, interface{}) error
	// 移除index位置的元素
	Remove(int) (interface{}, error)
	// 返回元素对应的索引
	IndexOf(interface{}) (int, error)
	// 清除所有元素
	Clear()
}
