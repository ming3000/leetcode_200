package string

func countBinarySubstrings(s string) int {
	preLen, curLen, count := 0, 1, 0
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			curLen++
		} else {
			preLen = curLen
			curLen = 1
		} // else>>

		if preLen >= curLen {
			count++
		} // if>
	} // for>

	return count
}
