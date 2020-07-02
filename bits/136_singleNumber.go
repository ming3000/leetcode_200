package bits

func singleNumber136(nums []int) int {
	var ret int
	for _, num := range nums {
		ret = ret ^ num
	} // for>
	return ret
}
