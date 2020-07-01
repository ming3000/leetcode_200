package hash

import "sort"

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	} // if>
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		} // if>
	} // for>
	return false
}
