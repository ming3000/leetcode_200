package bits

func hammingDistance(x int, y int) int {
	var count int
	z := x ^ y
	for z != 0 {
		if z&1 == 1 {
			count++
		} // if>>
		z = z >> 1
	} // for>
	return count
}
