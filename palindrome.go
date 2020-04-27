package main

func main() {
	print(isPalindrome(21))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}

	rev := 0
	nv := x
	for {
		rev += nv % 10
		nv /= 10
		if nv == 0 {
			break
		} else {
			rev *= 10
		}
	}

	return rev == x
}
