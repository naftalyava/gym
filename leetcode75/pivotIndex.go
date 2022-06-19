func runningSum(nums []int) []int {
	var sum int
	var sums []int

	for _, element := range nums {
		sum += element
		sums = append(sums, sum)
	}

	return sums
}

func pivotIndex(nums []int) int {

	if len(nums) == 0 {
		return -1
	}

	var new_nums []int
	new_nums = append(new_nums, 0)
	new_nums = append(new_nums, nums...)
	new_nums = append(new_nums, 0)

	sums := runningSum(new_nums)
	sum := sums[len(sums)-1]

	for i := 1; i < len(sums)-1; i++ {
		if sums[i-1] == sum-sums[i-1]-new_nums[i] {
			return i - 1
		}
	}

	return -1

}