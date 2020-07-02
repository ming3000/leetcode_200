package array

func matrixReshape(nums [][]int, r int, c int) [][]int {
	rawRows, rawColumns := len(nums), len(nums[0])
	if rawRows*rawColumns != r*c {
		return nums
	} // if>

	ans := make([][]int, r)
	for i := range ans {
		ans[i] = make([]int, c)
	} // for>

	i, j := 0, 0
	for _, row := range nums {
		for _, col := range row {
			ans[i][j] = col
			j++
			if j == c {
				i++
				j = 0
			} // if>>>
		} // for>>
	} // for>
	return ans
}
