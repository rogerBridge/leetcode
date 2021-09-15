package leetcode

func IsValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	m := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	stack := make([]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		if _, ok := m[string(s[i])]; !ok {
			stack = append(stack, string(s[i]))
		} else {
			if len(stack) == 0 {
				return false
			}
			if m[string(s[i])] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
