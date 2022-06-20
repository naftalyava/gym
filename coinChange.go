package main

import (
	"fmt"
	"math"
)

func coinChange_aux(coins []int, amount int, g_map *[10001]int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	found_in_cache := (*g_map)[amount]

	// we have something in cache
	if found_in_cache != -2 {
		return found_in_cache
	}

	// didnt found in cache, calculate
	min := math.MaxInt
	for _, e := range coins {
		tmp := coinChange_aux(coins, amount-e, g_map)

		if amount-e >= 0 {
			(*g_map)[amount-e] = tmp
		} else {
			continue
		}

		if tmp >= 0 && tmp <= min {
			min = tmp
		}
	}

	if min < math.MaxInt {
		return min + 1
	} else {
		return -1
	}
}

func coinChange(coins []int, amount int) int {
	var g_map [10001]int
	for i := range g_map {
		g_map[i] = -2
	}
	tmp := -1
	tmp = coinChange_aux(coins, amount, &g_map)
	return tmp
}

func main() {
	coins := []int{1, 2, 5}

	fmt.Println(coinChange(coins, 8))
}
