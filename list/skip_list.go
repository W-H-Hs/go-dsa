package list

import (
	"errors"
	"go-dsa/common"
	"math/rand"
)

// 跳表中的元素必须是有序的。
// 查找：给定元素key，从头结点的最顶层开始，检查最顶层的nextNode的key是否大于给定的key，如果大
// 		于，则检查头头结点第二层的nextNode的key是否大于给定的可以。如果nextNode的key小于给定
//		key，则开始检查nextNode最顶层的nextNode是否大于给定的key，如果大于则检查第二层，如果
//		小于则检查nextNode的nextNode。重复这个过程。

const (
	maxLevel = 32   // maxLevel 跳表的最高层数。
	p        = 0.25 // p 是计算跳表新增节点的level的概率阈值。
)

type skipListNode struct {
	key       common.ComparableElement // key 是跳表节点key-value键值对中的key。
	value     interface{}              // value 是跳表节点key-value键值对中的value。
	nextNodes []*skipListNode          // nextNodes 是跳表节点的next指针指向的所有下一节点。

	level int // level 是跳表节点的层级。
}

type SkipList struct {
	size  int // size 是跳表的元素个数。
	level int // level 是跳表的当前层数。

	header *skipListNode // header 是跳表的头指针。
}

// NewSkipList 构造函数返回SkipList结构体的指针。
func NewSkipList() *SkipList {
	return &SkipList{
		header: &skipListNode{nextNodes: make([]*skipListNode, maxLevel)},
	}
}

// Size 返回SkipList的元素数量。
func (s *SkipList) Size() int {
	return s.size
}

// IsEmpty 返回SkipList的元素是否为空。
func (s *SkipList) IsEmpty() bool {
	return s.size == 0
}

// randomLevel 使用随机函数计算新节点的层级。
func (s *SkipList) randomLevel() int {
	level := 0
	for rand.Float32() < p && level < maxLevel {
		level++
	}
	return level
}

// Put 向SkipList中添加key-value键值对并返回key原来对应的value。
func (s *SkipList) Put(key common.ComparableElement, value interface{}) (interface{}, error) {
	if key == nil || value == nil {
		return nil, errors.New("key must not be nil")
	}

	level := s.randomLevel()
	newNode := &skipListNode{
		key:       key,
		value:     value,
		nextNodes: make([]*skipListNode, maxLevel),
		level:     level,
	}
	prevNodes := make([]*skipListNode, level+1)
	// 将前驱节点全部初始化为header。
	for i := 0; i <= level; i++ {
		prevNodes[i] = s.header
	}

	// 寻找插入位置。
	cur := s.header
	for i := s.level; i >= 0; i-- {
		for {
			nextNode := cur.nextNodes[i]
			if nextNode == nil {
				prevNodes[i] = cur
				break
			}
			cmp := key.CompareTo(nextNode.key)
			if cmp == 0 {
				// 如果cmp等于0则代表key以存在于跳表中，直接更新value即可。
				oldValue := cur.value
				cur.value = value
				return oldValue, nil
			} else if cmp == 1 {
				// 如果cmp等于1，游标指针右移。
				cur = nextNode
			} else {
				// 如果cmp等于-1，游标指针下移。
				// 如果此时层级i小于新插入节点，则代表cur为新插入节点的前驱节点。
				if i <= level {
					prevNodes[i] = cur
				}
				break
			}
		}
	}

	// 插入节点。
	for idx, prevNode := range prevNodes {
		nextNode := prevNode.nextNodes[idx]
		newNode.nextNodes[idx] = nextNode
		prevNode.nextNodes[idx] = newNode
	}

	// 如果新节点的层数大于跳表的当前层数，则更新，在这一次插入中header的nextNodes数量会小于s.level，
	// 但在下一次插入的时候就会更新为s.level。
	if level > s.level {
		s.level = level
	}

	s.size++

	return nil, nil
}

// Get 获取SkipList中key对应的value。
func (s *SkipList) Get(key common.ComparableElement) (interface{}, error) {
	if key == nil {
		return nil, errors.New("key must not be nil")
	}

	cur := s.header
	for i := s.level; i >= 0; i-- {
		for {
			nextNode := cur.nextNodes[i]
			if nextNode == nil {
				break
			}
			cmp := key.CompareTo(nextNode.key)
			if cmp == 0 {
				// cmp等于0代表已经找到，直接返回即可。
				return cur.nextNodes[i].value, nil
			} else if cmp == 1 {
				// cmp等于1代表游标应该向右移。
				cur = nextNode
			} else {
				// cmp等于-1代表游标应该向下移。
				break
			}
		}
	}

	// 如果代码能够到达这里，代表跳表中不存在key。
	return nil, nil
}

// Remove 移除key对应的键值对，如果key存在于跳表中，则返回key对应的value，否则返回nil。
func (s *SkipList) Remove(key common.ComparableElement) (interface{}, error) {
	if key == nil {
		return nil, errors.New("key must not be nil")
	}

	var targetNode *skipListNode
	prevNodes := make([]*skipListNode, s.level+1)
	// 将前驱节点全部初始化为header。
	for i := 0; i <= s.level; i++ {
		prevNodes[i] = s.header
	}

	cur := s.header
	for i := s.level; i >= 0; i-- {
		for {
			nextNode := cur.nextNodes[i]
			if nextNode == nil {
				break
			}

			cmp := key.CompareTo(nextNode.key)
			if cmp == 0 {
				// cmp等于0，找到目标节点。
				targetNode = nextNode
				if i == 0 {
					// 如果i不等于0，则代表游标指针还应该向下移动来找到所有的前驱节点，否则退出循环。
					break
				}
			} else if cmp == 1 {
				// cmp等于1，游标指针右移。
				cur = nextNode
			} else {
				// cmp等于-1，游标指针下移，记录前驱节点。
				prevNodes[i] = cur
				break
			}
		}
	}

	if targetNode == nil {
		// 如果targetNode不存在则代表key不存在与跳表中。
		return nil, nil
	} else {
		for i := 0; i <= targetNode.level; i++ {
			prevNodes[i].nextNodes[i] = targetNode.nextNodes[i]
		}
		return targetNode.value, nil
	}
}
