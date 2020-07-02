package string

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	} // if>
	checkArr := [26]int{}
	for _, v := range s {
		checkArr[v-'a']++
	} // for>
	for _, v := range t {
		checkArr[v-'a']--
	} // for>

	for _, v := range checkArr {
		if v != 0 {
			return false
		}
	} // for>
	return true
}
