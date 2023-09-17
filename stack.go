package leetcode

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x, true
}

func (s *Stack) Peek() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	return (*s)[len(*s)-1], true
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}
