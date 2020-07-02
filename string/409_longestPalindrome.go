package string

//使用长度为 256 的整型数组来统计每个字符出现的个数，每个字符有偶数个可以用来构成回文字符串。
//因为回文字符串最中间的那个字符可以单独出现，所以如果有单独的字符就把它放到最中间。
func longestPalindrome(s string) int {
	checkArr := [256]int{}
	for i := 0; i < len(s); i++ {
		checkArr[s[i]-'a']++
	} // for>

	var palindrome int
	for _, v := range checkArr {
		palindrome += (v / 2) * 2
	} // for>
	if palindrome < len(s) {
		palindrome++
	}
	return palindrome
}
