package string

// 记录一个字符上次出现的位置，如果两个字符串中的字符上次出现的位置一样，那么就属于同构。
func isIsomorphic(s string, t string) bool {
	preIndexOfS := [256]int{}
	preIndexOfT := [256]int{}

	for i := 0; i < len(s); i++ {
		if preIndexOfS[s[i]] != preIndexOfT[t[i]] {
			return false
		} // if>
		preIndexOfS[s[i]] = i + 1
		preIndexOfT[t[i]] = i + 1
	} // for>

	return true
}
