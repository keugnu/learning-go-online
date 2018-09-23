package stack

import "testing"

var s = new(Stack)

func TestPush(t *testing.T) {
	s.Push(1)
	if s.Pop() == 1 {
		return
	} else {
		t.Log("Pop did not return 1.")
		t.Fail()
	}
}
