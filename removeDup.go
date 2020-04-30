package main

func main() {
	println(removeDuplicates([]int{1, 2, 2, 3}))
	println(removeDuplicates([]int{1, 2, 2}))
	println(removeDuplicates([]int{1, 2, 3}))
	println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	index := 0
	for i := 1; i < len(nums); {
		if nums[i] != nums[index] {
			index += 1
			if index != i {
				nums[index] = nums[i]
			}
		}
		i += 1
	}

	return index + 1
}
