package leetcode

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := rand.Intn(100) + 1
		firstBadVersionFirstBad = rand.Intn(n) + 1
		result := firstBadVersion(n)
		if result != firstBadVersionFirstBad {
			t.Errorf("firstBadVersion(%d) = %d, want %d, n = %d", n, result, firstBadVersionFirstBad, n)
		}
	}
}

func BenchmarkFirstBadVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Intn(100) + 1
		firstBadVersionFirstBad = rand.Intn(n) + 1
		result := firstBadVersion(n)
		if result != firstBadVersionFirstBad {
			fmt.Printf("%d != %d (n = %d)", result, firstBadVersionFirstBad, n)
		}
	}
}

func BenchmarkFirstBadVersionTopRatedV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Intn(100) + 1
		firstBadVersionFirstBad = rand.Intn(n) + 1
		result := firstBadVersionTopRatedV1(n)
		if result != firstBadVersionFirstBad {
			fmt.Printf("%d != %d (n = %d)", result, firstBadVersionFirstBad, n)
		}
	}
}

func BenchmarkFirstBadVersionTopRatedV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Intn(100) + 1
		firstBadVersionFirstBad = rand.Intn(n) + 1
		result := firstBadVersionTopRatedV2(n)
		if result != firstBadVersionFirstBad {
			fmt.Printf("%d != %d (n = %d)", result, firstBadVersionFirstBad, n)
		}
	}
}

func BenchmarkFirstBadVersionTopRatedV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Intn(100) + 1
		firstBadVersionFirstBad = rand.Intn(n) + 1
		result := firstBadVersionTopRatedV3(n)
		if result != firstBadVersionFirstBad {
			fmt.Printf("%d != %d (n = %d)", result, firstBadVersionFirstBad, n)
		}
	}
}

func BenchmarkFirstBadVersionTopRatedV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Intn(100) + 1
		firstBadVersionFirstBad = rand.Intn(n) + 1
		result := firstBadVersionTopRatedV4(n)
		if result != firstBadVersionFirstBad {
			fmt.Printf("%d != %d (n = %d)", result, firstBadVersionFirstBad, n)
		}
	}
}

func BenchmarkFirstBadVersionTopRatedV5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Intn(100) + 1
		firstBadVersionFirstBad = rand.Intn(n) + 1
		result := firstBadVersionTopRatedV5(n)
		if result != firstBadVersionFirstBad {
			fmt.Printf("%d != %d (n = %d)", result, firstBadVersionFirstBad, n)
		}
	}
}
