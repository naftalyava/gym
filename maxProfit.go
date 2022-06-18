package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	var max []int
	max = append(max, prices...)
	//var min []int = prices
	curMax := math.MinInt32
	//curMin := math.MaxInt32
	maxProfit := 0

	for i, _ := range prices {
		if prices[len(prices)-1-i] >= curMax {
			curMax = prices[len(prices)-1-i]
		}
		max[len(max)-1-i] = curMax
	}

	fmt.Println(max)
	fmt.Println(prices)

	for i, _ := range max {
		if max[i]-prices[i] > maxProfit {
			maxProfit = max[i] - prices[i]
		}
	}

	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices))
}

/*
{766664}
{715364}

*/
