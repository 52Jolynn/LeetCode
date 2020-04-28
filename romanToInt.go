package main

func main() {
	print(romanToInt("LVIII"))
}

//https://leetcode-cn.com/problems/roman-to-integer
func romanToInt(s string) int {
	roman := make(map[byte]int, 7)
	roman['I'] = 1
	roman['V'] = 5
	roman['X'] = 10
	roman['L'] = 50
	roman['C'] = 100
	roman['D'] = 500
	roman['M'] = 1000

	b := []byte(s)
	length := len(b)

	//通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
	//
	//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
	//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
	//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

	idx := length - 1
	ret := 0
	for {
		bb := b[idx]
		ret += roman[bb]
		if idx == 0 {
			break
		}
		idx -= 1
		if idx >= 0 {
			//往前读多一位
			nbb := b[idx]
			if (nbb == 'I' && (bb == 'V' || bb == 'X')) || (nbb == 'X' && (bb == 'L' || bb == 'C')) || (nbb == 'C' && (bb == 'D' || bb == 'M')) {
				ret -= roman[nbb]
				idx -= 1
			}
		}
		if idx < 0 {
			break
		}
	}

	return ret
}
