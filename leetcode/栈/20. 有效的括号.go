package 栈

// https://leetcode-cn.com/problems/valid-parentheses/

type stack struct {
	array []byte
}

func (s *stack) push(ele byte) {
	s.array = append(s.array, ele)
}
func (s *stack) pop() byte {
	ele := s.array[len(s.array)-1]
	newArray := make([]byte, len(s.array)-1)
	copy(newArray, s.array[:len(s.array)-1])
	s.array = newArray
	return ele
}
func (s *stack) len() int {
	return len(s.array)
}

func isValid(s string) bool {
	if s == "" {
		return true
	}

	if len(s)&1 == 1 {
		return false
	}

	const (
		small = "()"
		mid   = "[]"
		big   = "{}"
	)

	stack := stack{}

	for i := 0; i < len(s); i++ {
		if s[i] == small[0] || s[i] == mid[0] || s[i] == big[0] {
			stack.push(s[i])
			continue
		}
		if stack.len() != 0 {
			ele := stack.pop()
			if s[i] == small[1] && string(ele)+string(s[i]) != small {
				return false
			}
			if s[i] == mid[1] && string(ele)+string(s[i]) != mid {
				return false
			}
			if s[i] == big[1] && string(ele)+string(s[i]) != big {
				return false
			}
		} else {
			return false
		}
	}
	if stack.len() != 0 {
		return false
	}
	return true
}
