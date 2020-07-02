package string

// 从字符串的某一位开始，尝试着去扩展子字符串。
func countSubstrings(s string) int {
	var count int
	for i := 0; i < len(s); i++ {
		// 偶数长度
		extendSubStr(s, i, i, &count)
		// 奇数长度
		extendSubStr(s, i, i+1, &count)
	}
	return count
}

func extendSubStr(s string, start, end int, count *int) {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
		*count += 1
	} // for>
}
