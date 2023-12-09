package leetcode

import (
	"testing"
)

func TestImplementQueueUsingStacks(t *testing.T) {
	queue := ConstructorTopRated()
	queue.Push(1)
	queue.Push(2)
	x := queue.Peek()
	if x != 1 {
		t.Errorf("queue.Peek() returned %d, want 1", x)
	}
	x = queue.Pop()
	if x != 1 {
		t.Errorf("queue.Pop() returned %d, want 1", x)
	}
	res := queue.Empty()
	if res != false {
		t.Errorf("queue.Empty() returned %t, want false", res)
	}
}
