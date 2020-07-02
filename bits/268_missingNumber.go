package bits

func missingNumber(nums []int) int {
	var ret int
	for i := 0; i < len(nums); i++ {
		ret = ret ^ i ^ nums[i]
	} // for>
	return ret ^ len(nums)
}
