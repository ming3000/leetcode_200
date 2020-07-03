package bst

import "leetcode_200/common"

func findTarget(root *common.TreeNode, k int) bool {
	nums := make([]int, 0)
	inOrder(root, &nums)

	checkMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := k - nums[i]
		if _, exists := checkMap[diff]; exists {
			return true
		} else {
			checkMap[nums[i]] = i
		} // else>>
	} // for>
	return false
}

func inOrder(root *common.TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inOrder(root.Right, nums)
}
