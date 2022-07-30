package solution

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func MaxSubArray(nums []int) int {
	localMax := nums[0]
	globalMax := nums[0]

	for i := 1; i < len(nums); i++ {
		localMax = Max(localMax+nums[i], nums[i])
		globalMax = Max(globalMax, localMax)
	}

	return globalMax
}
