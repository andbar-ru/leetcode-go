package leetcode

import (
	"fmt"
	"math/rand"
	"time"
)

func randSliceInt(minLen, maxLen int, minN, maxN int, seed int64) ([]int, error) {
	if minLen > maxLen {
		return nil, fmt.Errorf("minLen > maxLen (%d > %d)", minLen, maxLen)
	}
	if minN > maxN {
		return nil, fmt.Errorf("minN > maxN (%d > %d)", minN, maxN)
	}

	if seed == -1 {
		seed = time.Now().UnixNano()
	}
	r := rand.New(rand.NewSource(seed))
	l := r.Intn(maxLen+1-minLen) + minLen
	slice := make([]int, l)
	for i := 0; i < l; i++ {
		slice[i] = r.Intn(maxN+1-minN) + minN
	}

	return slice, nil
}
