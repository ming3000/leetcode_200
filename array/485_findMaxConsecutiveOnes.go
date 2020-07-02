package array

import "leetcode_200/common"

func findMaxConsecutiveOnes(nums []int) int {
	var cur, max int
	for _, num := range nums {
		if num == 0 {
			cur = 0
		} else {
			cur++
		} // else>>
		max = common.Max(max, cur)
	} // for>
	return max
}
