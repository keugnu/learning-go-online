package stack

type Stack struct {
	ptr int
	stk [100]int
}

func (s *Stack) Push(n int) {
	if s.ptr+1 > len(s.stk) {
		return
	} else {
		s.ptr++
		s.stk[s.ptr] = n
		return
	}
}

func (s *Stack) Pop() (n int) {
	n = s.stk[s.ptr]
	s.ptr--
	return
}
