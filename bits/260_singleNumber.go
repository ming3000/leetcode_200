package bits

//首先计算 bitmask ^= x，则 bitmask 不会保留出现两次数字的值，因为相同数字的异或值为 0。
//但是 bitmask 会保留只出现一次的两个数字（x 和 y）之间的差异。
//我们可以直接从 bitmask 中提取 x 和 y 吗？不能，但是我们可以用 bitmask 作为标记来分离 x 和 y。
//我们通过 bitmask & (-bitmask) 保留 bitmask 最右边的 1，这个 1 要么来自 x，要么来自 y。
//当我们找到了 x，那么 y = bitmask^x

func singleNumber260(nums []int) []int {
	// 将数组的所有元素异或得到的结果为不存在重复的两个元素异或的结果
	var diff int
	for _, num := range nums {
		diff ^= num
	} // for>
	// x & (-x) 是保留位中最右边 1 ，且将其余的 1 设位 0 的方法
	diff &= -diff
	ret := make([]int, 2)
	for _, num := range nums {
		if num&diff == 0 {
			ret[0] ^= num
		} else {
			ret[1] ^= num
		} // else>>
	} // for>
	return ret
}
