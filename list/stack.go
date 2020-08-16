package list

type Stack struct {
	array []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Top() interface{} {
	if len(s.array) == 0 {
		return 0
	}
	return s.array[len(s.array)-1]
}

func (s *Stack) Push(ele interface{}) {
	s.array = append(s.array, ele)
}

func (s *Stack) Pop() interface{} {
	ele := s.array[len(s.array)-1]
	newArray := make([]interface{}, len(s.array)-1)
	copy(newArray, s.array[:len(s.array)-1])
	s.array = newArray
	return ele
}

func (s *Stack) Size() int {
	return len(s.array)
}
