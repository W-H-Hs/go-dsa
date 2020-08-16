package 二叉树

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	q := queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	fmt.Println(q.dequeue())
	fmt.Println(q.len())
	fmt.Println(q.dequeue())
	fmt.Println(q.len())
	fmt.Println(q.dequeue())
	fmt.Println(q.len())
}
