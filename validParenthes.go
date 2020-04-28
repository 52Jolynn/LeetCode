package main

func main() {
	print(isValid("([)"))    //false
	print(isValid("()[]{}")) //true
	print(isValid("([)]"))   //false
	print(isValid("{[]}"))   //true
	print(isValid("()"))     //true
	print(isValid("(]"))     //false
	print(isValid("}{"))     //false
}

//https://leetcode-cn.com/problems/valid-parentheses
func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	var stack []byte
	top := -1
	for i, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, byte(v))
			top += 1
		} else if v == ')' || v == '}' || v == ']' {
			if top < 0 {
				return false
			}
			topEle := stack[top]
			if topEle == '(' && topEle+1 != byte(v) {
				return false
			} else if (topEle == '{' || topEle == '[') && topEle+2 != byte(v) {
				return false
			}

			top -= 1
			if top < 0 {
				stack = nil
			} else {
				stack = stack[:top+1]
			}
		} else {
			return false
		}
		if i >= len(s) {
			break
		}
	}
	return len(stack) == 0
}
