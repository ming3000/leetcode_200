package recursion

import (
	"leetcode_200/common"
	"math"
)

func isBalanced(root *common.TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if math.Abs(leftHeight-rightHeight) > 1 {
		return false
	} // if>

	return isBalanced(root.Left) && isBalanced(root.Right)
}

//计算节点最大深度
func height(node *common.TreeNode) float64 {
	if node == nil {
		return 0
	}
	return math.Max(height(node.Left), height(node.Right)) + 1
}
