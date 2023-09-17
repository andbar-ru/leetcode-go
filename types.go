package leetcode

import "fmt"

type Int struct {
	v int
}

func (i *Int) String() string {
	return fmt.Sprint(i.v)
}
