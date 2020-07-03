package bfs

import "leetcode_200/common"

func findBottomLeftValue(root *common.TreeNode) int {
	if root == nil {
		return 0
	} // if>

	ret := make([][]int, 0)
	queue := make([]*common.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		curLevelLen := len(queue)
		curLevelList := make([]int, 0)
		for i := 0; i < curLevelLen; i++ {
			head := queue[0]
			queue = queue[1:]

			if head.Left != nil {
				queue = append(queue, head.Left)
			} // if>>>
			if head.Right != nil {
				queue = append(queue, head.Right)
			} //if>>>
			curLevelList = append(curLevelList, head.Val)
		} // for>>
		ret = append(ret, curLevelList)
	} // for>

	return ret[len(ret)-1][0]
}
