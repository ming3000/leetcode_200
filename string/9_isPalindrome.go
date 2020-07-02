package string

func isPalindrome(x int) bool {
	// 负数肯定不是回文，因为负号
	if x < 0 {
		return false
	} // if>

	var rawNum = x
	var reversedNum int
	for rawNum != 0 {
		reversedNum = reversedNum*10 + rawNum%10
		rawNum = rawNum / 10
	} // for>
	return x == reversedNum
}
