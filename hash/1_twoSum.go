package hash

func twoSum(nums []int, target int) []int {
	if len(nums) < 0 {
		return []int{-1, -1}
	} // if>
	// key - the num, value - the index
	cm := make(map[int]int, 0)
	for i := range nums {
		diff := target - nums[i]
		if index, ok := cm[diff]; ok {
			return []int{i, index}
		} else {
			cm[nums[i]] = i
		} // if>
	} // for>

	return []int{-1, -1}
}
