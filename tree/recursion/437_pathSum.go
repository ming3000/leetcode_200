package recursion

import "leetcode_200/common"

func pathSum(root *common.TreeNode, sum int) int {
	if root == nil {
		return 0
	} // if>
	ret := pathSumStartWithRoot(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
	return ret
}

func pathSumStartWithRoot(root *common.TreeNode, sum int) int {
	if root == nil {
		return 0
	} // if>

	var ret int
	if root.Val == sum {
		ret++
	} // if>

	ret += pathSumStartWithRoot(root.Left, sum-root.Val) + pathSumStartWithRoot(root.Right, sum-root.Val)
	return ret
}
