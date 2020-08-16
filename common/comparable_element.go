package common

type ComparableElement interface {
	// 返回-1，表示self小于参数
	// 返回0，表示self等于参数
	// 返回1，表示self大于参数
	CompareTo(ComparableElement) int
}

type Int int

func (c Int) CompareTo(ele ComparableElement) int {
	if c > (ele).(Int) {
		return 1
	}
	if c < (ele).(Int) {
		return -1
	}
	return 0
}
