package main

import "bytes"

func main() {
	data := make([]string, 3)
	data[0] = "flower"
	data[1] = "flow"
	data[2] = "flight"
	print(longestCommonPrefix(data))

	data = make([]string, 1)
	data[0] = ""
	print(longestCommonPrefix(data))
}

//https://leetcode-cn.com/problems/longest-common-prefix
func longestCommonPrefix(strs []string) string {
	strLen := len(strs)
	if strLen == 0 {
		return ""
	}

	var prefix bytes.Buffer
	cur := byte('-')
	idx := 0
	endless := false
	for {
		for i, v := range strs {
			length := len(v)
			if idx > length-1 {
				endless = true
				break
			}
			if cur == '-' {
				cur = v[idx]
			} else if cur != v[idx] {
				endless = true
				break
			}
			if i == strLen-1 {
				idx += 1
				break
			}
		}
		if endless {
			break
		}
		prefix.WriteByte(cur)
		cur = '-'
	}

	return prefix.String()
}
