package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	hash := make(map[int]int)

	for i, e := range nums {

		old_i, found := hash[target-e]
		if found {
			return []int{old_i, i}
		}

		hash[e] = i
	}

	return []int{-1, -1}
}

func main() {
	var nums = []int{3, 1, 3}
	target := 6

	a := twoSum(nums, target)

	fmt.Println(a)

}
