package hash

import "leetcode_200/common"

func findLHS(nums []int) int {
	checkMap := make(map[int]int)
	for _, num := range nums {
		if value, exists := checkMap[num]; exists {
			checkMap[num] = value + 1
		} else {
			checkMap[num] = 1
		} // else>>
	} // for>

	longest := 0
	for k, _ := range checkMap {
		if _, exists := checkMap[k+1]; exists {
			longest = common.Max(longest, checkMap[k]+checkMap[k+1])
		} // if>>
	} // for>

	return longest
}
