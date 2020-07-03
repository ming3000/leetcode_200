package bst

import "leetcode_200/common"

func sortedArrayToBST(nums []int) *common.TreeNode {
	if len(nums) == 0 {
		return nil
	} // if>

	mid := len(nums) / 2
	root := &common.TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
