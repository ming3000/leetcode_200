package bst

import (
	"leetcode_200/common"
	"math"
)

func getMinimumDifference(root *common.TreeNode) int {
	inOrderPath := make([]int, 0)
	inOrderTraversal(root, &inOrderPath)
	minDiff := math.MaxInt64
	for i := 1; i < len(inOrderPath); i++ {
		diff := common.Abs(inOrderPath[i] - inOrderPath[i-1])
		if diff < minDiff {
			minDiff = diff
		} // if>>
	} // for>

	return minDiff
}

func inOrderTraversal(root *common.TreeNode, nums *[]int) {
	if root == nil {
		return
	} // if>
	inOrderTraversal(root.Left, nums)
	*nums = append(*nums, root.Val)
	inOrderTraversal(root.Right, nums)
}
