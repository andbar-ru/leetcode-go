package leetcode

import (
	"fmt"
	"testing"
)

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	var tests = []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
	}

	for _, test := range tests {
		result := bestTimeToBuyAndSellStock(test.prices)
		if result != test.want {
			t.Errorf("bestTimeToBuyAndSellStock(%v) = %d, want %d", test.prices, result, test.want)
		}
	}
}

func bestTimeToBuyAndSellStockBrute(prices []int) int {
	var maxProfit int

	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

func TestBestTimeToBuyAndSellStockRand(t *testing.T) {
	numTests := 1000

	for i := 0; i < numTests; i++ {
		if i%10 == 0 {
			fmt.Printf("%d", i)
		} else {
			fmt.Print(".")
		}

		seed := int64(i)
		prices, err := randSliceInt(1, 1e5, 0, 1e4, seed)
		if err != nil {
			panic(err)
		}
		result := bestTimeToBuyAndSellStock(prices)
		want := bestTimeToBuyAndSellStockBrute(prices)
		if result != want {
			t.Errorf("bestTimeToBuyAndSellStock(%v) = %d, want %d", prices, result, want)
			break
		}
	}
}
