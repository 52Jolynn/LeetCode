package main

func main() {
	nums := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(nums)
	for i := 0; i < len(nums); i++ {
		print(nums[i])
	}
}

//https://leetcode-cn.com/problems/duplicate-zeros/
func duplicateZeros(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		if arr[i] == 0 {
			if i+1 >= length {
				break
			}
			//平移
			for j := length - 1; j > i+1; j-- {
				arr[j] = arr[j-1]
			}
			i += 1
			arr[i] = 0
		}
	}
}
