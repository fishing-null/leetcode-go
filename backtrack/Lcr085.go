package backtrack

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	generateParenthesisBacktrack(n, n, "", &ret)
	return ret
}

func generateParenthesisBacktrack(left, right int, track string, result *[]string) {

	if right < left {
		return
	}

	if left == 0 && right == 0 {
		*result = append(*result, track)
		return
	}

	if left < 0 || right < 0 {
		return
	}

	generateParenthesisBacktrack(left-1, right, track+"(", result)
	generateParenthesisBacktrack(left, right-1, track+")", result)
}

func generateParenthesisIsValid(s string) bool {
	stack := make([]rune, 0)
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			// 左括号入栈
			stack = append(stack, char)
		case ')', ']', '}':
			// 右括号需要匹配栈顶元素
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			// 匹配成功，弹出栈顶元素
			stack = stack[:len(stack)-1]
		}
	}

	// 最后栈应该为空才表示完全匹配
	return len(stack) == 0
}
