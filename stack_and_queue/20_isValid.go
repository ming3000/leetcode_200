package stack_and_queue

func isValid(s string) bool {
	if s == "" {
		return true
	} // if>

	checkMap := map[byte]byte{')': '(', ']': '[', '}': '{'}
	checkStack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			checkStack = append(checkStack, s[i])
		} else {
			if len(checkStack) == 0 {
				return false
			} // if>>>
			if checkMap[s[i]] != checkStack[len(checkStack)-1] {
				return false
			} // if>>>
			checkStack = checkStack[:len(checkStack)-1]
		} // else>>
	} // for>

	return len(checkStack) == 0
}
