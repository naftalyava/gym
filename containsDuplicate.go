package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	var hash = make(map[int]bool)

	for _, e := range nums {
		_, found := hash[e]
		if found {
			return true
		} else {
			hash[e] = true
		}
	}

	return false
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(containsDuplicate(nums))
}
