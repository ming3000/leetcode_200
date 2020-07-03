package two_pointer

//使用双指针可以很容易判断一个字符串是否是回文字符串：
//令一个指针从左到右遍历，一个指针从右到左遍历，
//这两个指针同时移动一个位置，每次都判断两个指针指向的字符是否相同，
//如果都相同，字符串才是具有左右对称性质的回文字符串。
//
//本题的关键是处理删除一个字符。
//在使用双指针遍历字符串时，如果出现两个指针指向的字符不相等的情况，我们就试着删除一个字符，再判断删除完之后的字符串是否是回文字符串。
//在判断是否为回文字符串时，我们不需要判断整个字符串，
//因为左指针左边和右指针右边的字符之前已经判断过具有对称性质，所以只需要判断中间的子字符串即可。
//在试着删除字符时，我们既可以删除左指针指向的字符，也可以删除右指针指向的字符。
func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s, left, right-1) || isPalindrome(s, left+1, right)
		} // if>>
		left++
		right--
	} // for>
	return true
}

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		} // if>>
		start++
		end--
	} // for>
	return true
}
