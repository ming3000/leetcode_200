package bfs

import "leetcode_200/common"

func averageOfLevels(root *common.TreeNode) []float64 {
	if root == nil {
		return nil
	}

	ret := make([]float64, 0)
	queue := make([]*common.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		currentLevelLength := len(queue)
		currentLevelValueList := make([]int, 0)
		for i := 0; i < currentLevelLength; i++ {
			head := queue[0]
			queue = queue[1:]
			currentLevelValueList = append(currentLevelValueList, head.Val)
			if head.Left != nil {
				queue = append(queue, head.Left)
			} // if>>>
			if head.Right != nil {
				queue = append(queue, head.Right)
			} // if>>>
		} // for>>
		ret = append(ret, averageValue(currentLevelValueList))
	} // for>

	return ret
}

func averageValue(nums []int) float64 {
	var sum int
	for _, v := range nums {
		sum += v
	} // for>
	return float64(sum) / float64(len(nums))
}
