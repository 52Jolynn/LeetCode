package main

func main() {
	nums := []int{1, 0, 1}
	moveZeroes(nums)
	for i := 0; i < len(nums); i++ {
		print(nums[i])
	}
}

func moveZeroes(nums []int) {
	swap := 1
	length := len(nums)
	for i := 0; i < length && swap < length; {
		if nums[i] != 0 {
			i += 1
		} else if nums[i] == 0 && nums[swap] != 0 {
			nums[i] = nums[swap]
			nums[swap] = 0
			i += 1
		}
		swap += 1
	}
}
