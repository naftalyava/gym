func runningSum(nums []int) []int {
	var sum int
	var sums []int

	for _, element := range nums {
		sum += element
		sums = append(sums, sum)
	}

	return sums
}