package solution

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if j, exists := m[num]; exists {
			return []int{j, i}
		}

		m[target-num] = i
	}

	return []int{0, 0}
}
