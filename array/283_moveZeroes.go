package array

func moveZeroes(nums []int) {
	var index int
	for _, v := range nums {
		if v != 0 {
			nums[index] = v
			index++
		} // if>>
	} // for>

	for index < len(nums) {
		nums[index] = 0
		index++
	} // for>
}
