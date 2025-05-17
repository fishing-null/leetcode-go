package backtrack

import "strings"

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	generateParenthesisBacktrack(n, 0, "", &ret)
	return ret
}

func generateParenthesisBacktrack(n, i int, track string, result *[]string) {
	if i == 2*n {
		if generateParenthesisIsValid(track) {
			*result = append(*result, track)
		}
		return
	}

	for _, ch := range [2]rune{'(', ')'} {
		sb := strings.Builder{}
		sb.WriteRune(ch)
		sb.WriteString(track)
		track = sb.String()
		generateParenthesisBacktrack(n, i+1, sb.String(), result)
		track = track[1:]
	}
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
